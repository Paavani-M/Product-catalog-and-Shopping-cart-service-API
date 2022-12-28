package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	post := typedefs.Product_Master{}

	err := json.Unmarshal(reqBody, &post)
	if err != nil {
		fmt.Println(err)
	}

	response := typedefs.Json_Response{}

	jsonStr, _ := json.Marshal(post.Specification)

	db := helpers.SetupDB()

	fmt.Println("Inserting details into DB")

	fmt.Println("Inserting product name  :" + post.Name)
	var price_round_off = math.Floor(float64(post.Price)*100) / 100
	_, err = db.Exec("INSERT INTO product_master(product_id,name,specification,sku,category_id,price) VALUES($1,$2,$3,$4,$5,$6);", post.Product_id, post.Name, string(jsonStr), post.Sku, post.Category_id, price_round_off)

	// check errors
	helpers.CheckErr(err)

	response = typedefs.Json_Response{Type: "success", Message: "Record has been inserted successfully!"}
	fmt.Println("Your data has been inserted successfuly")

	json.NewEncoder(w).Encode(response)

}
