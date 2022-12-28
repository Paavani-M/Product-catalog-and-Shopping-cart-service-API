package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func InsertCategory(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	category := typedefs.Category_Master{}

	err := json.Unmarshal(reqBody, &category)
	if err != nil {
		fmt.Println(err)
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	fmt.Println("Inserting details into DB")

	fmt.Println("Inserting category of name:  " + category.Name)

	_, err = db.Exec("INSERT INTO category_master(category_id,name) VALUES($1,$2);", category.Category_id, category.Name)
	helpers.CheckErr(err)

	response = typedefs.Json_Response{Type: "success", Message: "Record has been inserted successfully!"}
	fmt.Println("Your data has been inserted successfuly")

	json.NewEncoder(w).Encode(response)

}
