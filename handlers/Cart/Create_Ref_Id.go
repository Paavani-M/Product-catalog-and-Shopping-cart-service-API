package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"task.com/helpers"
	"task.com/typedefs"
)

func CreateRefId(w http.ResponseWriter, r *http.Request) {
	var response = typedefs.Json_Response{}

	id := uuid.New()
	db := helpers.SetupDB()
	currentdate := time.Now()

	_, err := db.Exec(helpers.CreateRefId, id.String(), currentdate)
	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	response = typedefs.Json_Response{Type: helpers.Success, Message: helpers.InsertSuccess}

	json.NewEncoder(w).Encode(response)

}
