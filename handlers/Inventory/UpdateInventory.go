package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func UpdateInventory(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	inventory := typedefs.Inventory{}

	err := json.Unmarshal(reqBody, &inventory)
	if err != nil {
		fmt.Println(err)
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	result, err := db.Exec("UPDATE inventory SET quantity=$1 WHERE product_Id=$2;", inventory.Quantity, inventory.Product_id)

	// check errors
	helpers.CheckErr(err)

	rows, err := result.RowsAffected()

	if rows != 1 {
		response = typedefs.Json_Response{Type: "missing", Message: "Inventory id doesn't exist"}
	} else {
		fmt.Println("Updating DB")
		fmt.Println("Updating product id:", inventory.Product_id)
		response = typedefs.Json_Response{Type: "success", Message: "Database has been updated successfully!"}
	}

	json.NewEncoder(w).Encode(response)

}
