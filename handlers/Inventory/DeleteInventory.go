package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"task.com/helpers"
	"task.com/typedefs"
)

func DeleteInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	Id := params["id"]

	a, _ := strconv.Atoi(Id)

	if a <= 0 {
		helpers.SendErrResponse(helpers.Error, helpers.ValidInput, w)
		return
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	result, err := db.Exec(helpers.DeleteInventory, Id)
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
		response = typedefs.Json_Response{Type: helpers.Success, Message: helpers.DeleteSuccess}
	}

	json.NewEncoder(w).Encode(response)

}
