package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("inside delete inventory function")
	Id := params["id"]

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	result, err := db.Exec("DELETE FROM inventory where product_id = $1", Id)

	helpers.CheckErr(err)

	rows, err := result.RowsAffected()

	helpers.CheckErr(err)

	if rows != 1 {
		response = typedefs.Json_Response{Type: "missing", Message: "Inventory id doesn't exist"}
	} else {
		fmt.Println("Deleting")
		response = typedefs.Json_Response{Type: "success", Message: "Deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)

}
