package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func InsertInventory(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	post := typedefs.Inventory{}

	err := json.Unmarshal(reqBody, &post)
	if err != nil {
		fmt.Println(err)
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	fmt.Println("Inserting...")

	_, err = db.Exec("INSERT INTO inventory(product_id,quantity) VALUES($1,$2);", post.Product_id, post.Quantity)
	helpers.CheckErr(err)

	response = typedefs.Json_Response{Type: "success", Message: "Record has been inserted successfully!"}
	fmt.Println("Your data has been inserted successfuly")

	json.NewEncoder(w).Encode(response)

}
