package handlers

import (
	"encoding/json"
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
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Unmarshalling, w)
		helpers.LogError(err)
		return
	}

	if inventory.Product_id <= 0 || inventory.Quantity <= 0 {
		helpers.SendErrResponse(helpers.Error, helpers.ValidInput, w)
		return
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.UpdateInventory, inventory.Quantity, inventory.Product_id)

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

	if rows != 1 {
		response = typedefs.Json_Response{Type: helpers.Missing, Message: helpers.Idnotexits}
	} else {
		response = typedefs.Json_Response{Type: helpers.Success, Message: helpers.UpdateSuccess}
	}

	json.NewEncoder(w).Encode(response)

}
