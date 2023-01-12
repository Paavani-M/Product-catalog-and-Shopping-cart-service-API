package handlers

import (
	"encoding/json"
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
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Unmarshalling, w)
		helpers.LogError(err)
		return
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	rows, err := db.Query(helpers.GetProductAll, product.Product_id)
	defer rows.Close()

	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

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

	a, err := db.Exec(helpers.UpdateProduct, product.Name, product.Product_id, string(jsonStr), product.Sku, product.Category_id, product.Price)

	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	b, err := a.RowsAffected()
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if b != 1 {
		response = typedefs.Json_Response{Type: helpers.Missing, Message: helpers.Idnotexits}
	} else {
		response = typedefs.Json_Response{Type: helpers.Success, Message: helpers.UpdateSuccess}

	}

	json.NewEncoder(w).Encode(response)
}
