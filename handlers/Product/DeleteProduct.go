package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("inside delete product function")
	Id := params["id"]

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()
	result, err := db.Exec("DELETE FROM product_master where product_id = $1", Id)

	helpers.CheckErr(err)

	rows, err := result.RowsAffected()

	helpers.CheckErr(err)

	if rows != 1 {
		response = typedefs.Json_Response{Type: "missing", Message: "Product id doesn't exist"}
	} else {
		fmt.Println("Deleting a product from DB")
		response = typedefs.Json_Response{Type: "success", Message: "Product has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)

}
