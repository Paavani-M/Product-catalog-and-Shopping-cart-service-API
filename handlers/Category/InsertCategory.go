package handlers

import (
	"encoding/json"
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
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Unmarshalling, w)
		helpers.LogError(err)
		return
	}

	if category.Category_id <= 0 {
		helpers.SendErrResponse(helpers.Error, helpers.ValidInput, w)
		return
	}

	var response = typedefs.Json_Response{}

	db := helpers.SetupDB()

	_, err = db.Exec(helpers.InsertCategory, category.Category_id, category.Name)
	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	response = typedefs.Json_Response{Type: helpers.Success, Message: helpers.InsertSuccess}

	json.NewEncoder(w).Encode(response)

}
