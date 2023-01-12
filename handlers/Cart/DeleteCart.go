package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	delete := typedefs.Delete_Cart{}

	err := json.Unmarshal(reqBody, &delete)
	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Unmarshalling, w)
		helpers.LogError(err)
		return
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.DeleteCart, delete.Reference_Id, delete.Product_Id)

	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if rows != 1 {
		response = typedefs.Json_Response{Type: helpers.Missing, Message: helpers.CartIdnotexists}
	} else {
		response = typedefs.Json_Response{Type: helpers.Success, Message: helpers.DeleteSuccess}
	}
	json.NewEncoder(w).Encode(response)

}
