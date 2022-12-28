package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	product := typedefs.Product_Master{}

	err := json.Unmarshal(reqBody, &product)
	if err != nil {
		fmt.Println(err)
	}
	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	rows, err := db.Query("SELECT * FROM product_master where product_id=$1", product.Product_id)
	defer rows.Close()

	result := typedefs.Product_Master{}

	var spec []byte

	for rows.Next() {
		rows.Scan(&result.Product_id, &result.Name, &spec, &result.Sku, &result.Category_id, &result.Price)
	}

	json.Unmarshal(spec, &result.Specification)

	var jsonStr []byte

	if product.Name == "" {
		product.Name = result.Name
	}

	if len(product.Specification) != 0 {
		jsonStr, _ = json.Marshal(product.Specification)
	} else {
		jsonStr, _ = json.Marshal(result.Specification)
	}

	if product.Sku == "" {
		product.Sku = result.Sku
	}

	if product.Category_id == 0 {
		product.Category_id = result.Category_id
	}

	if product.Price == 0 {
		product.Price = result.Price
	}

	a, err := db.Exec("UPDATE product_master SET Name=$1, Specification=$3, sku=$4, category_id=$5, price=$6 WHERE Product_id=$2;", product.Name, product.Product_id, string(jsonStr), product.Sku, product.Category_id, product.Price)
	// check errors
	helpers.CheckErr(err)

	b, err := a.RowsAffected()

	if b != 1 {
		response = typedefs.Json_Response{Type: "missing", Message: "product id doesn't exists!"}
	} else {
		fmt.Println("Updating DB")
		response = typedefs.Json_Response{Type: "success", Message: "database has been updated successfully!"}

	}

	json.NewEncoder(w).Encode(response)
}
