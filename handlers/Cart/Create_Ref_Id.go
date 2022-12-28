package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"task.com/helpers"
	"task.com/typedefs"
)

func CreateRefId(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside creating reference id for user function")
	var response = typedefs.Json_Response{}

	id := uuid.New()
	db := helpers.SetupDB()
	currentdate := time.Now()

	_, err := db.Exec("INSERT INTO cart_reference(reference_id,created_at) VALUES($1,$2);", id.String(), currentdate)
	helpers.CheckErr(err)

	response = typedefs.Json_Response{Type: "success", Message: "Record has been inserted successfully!"}
	fmt.Println("Your data has been inserted successfuly")

	json.NewEncoder(w).Encode(response)

}
