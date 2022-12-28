package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("inside delete category function")
	Id := params["id"]

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	fmt.Println("Deleting a category from DB")

	result, err := db.Exec("DELETE FROM category_master where category_id = $1", Id)

	helpers.CheckErr(err)

	rows, err := result.RowsAffected()

	helpers.CheckErr(err)

	if rows != 1 {
		response = typedefs.Json_Response{Type: "missing", Message: "Category id doesn't exist"}
	} else {
		response = typedefs.Json_Response{Type: "success", Message: "Category type has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)

}
