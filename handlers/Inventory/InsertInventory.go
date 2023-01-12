package handlers

import (
	"encoding/json"
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
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Unmarshalling, w)
		helpers.LogError(err)
		return
	}

	if post.Product_id <= 0 || post.Quantity <= 0 {
		helpers.SendErrResponse(helpers.Error, helpers.ValidInput, w)
		return
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.InsertUpdateInventory, post.Product_id)
	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	rows, err := result.RowsAffected()
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if rows == 1 {
		response = typedefs.Json_Response{Type: helpers.Update, Message: helpers.InsertUpdateInven}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = db.Exec(helpers.InsertInventory, post.Product_id, post.Quantity)
	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	response = typedefs.Json_Response{Type: helpers.Success, Message: helpers.InsertSuccess}

	json.NewEncoder(w).Encode(response)

}
