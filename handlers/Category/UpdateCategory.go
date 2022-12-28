package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func UpdateCategory(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	category := typedefs.Category_Master{}

	err := json.Unmarshal(reqBody, &category)
	if err != nil {
		fmt.Println(err)
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	result, err := db.Exec("UPDATE category_master SET Name=$1 WHERE Category_Id=$2;", category.Name, category.Category_id)

	// check errors
	helpers.CheckErr(err)

	rows, err := result.RowsAffected()

	if rows != 1 {
		response = typedefs.Json_Response{Type: "missing", Message: "Category id doesn't exist"}
	} else {
		fmt.Println("Updating DB")
		fmt.Println("Updating category id:", category.Category_id)
		response = typedefs.Json_Response{Type: "success", Message: "database has been updated successfully!"}
	}

	json.NewEncoder(w).Encode(response)

}
