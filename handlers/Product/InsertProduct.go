package handlers

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"

	"task.com/helpers"
	"task.com/typedefs"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	post := typedefs.Product_Master{}

	response := typedefs.Json_Response{}

	err := json.Unmarshal(reqBody, &post)
	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Unmarshalling, w)
		helpers.LogError(err)
		return
	}

	if post.Product_id <= 0 || post.Category_id <= 0 || post.Price <= 0 {
		helpers.SendErrResponse(helpers.Error, helpers.ValidInput, w)
		return
	}

	jsonStr, err := json.Marshal(post.Specification)
	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Marshalling, w)
		helpers.LogError(err)
		return
	}

	db := helpers.SetupDB()

	var price_round_off = math.Floor(float64(post.Price)*100) / 100

	_, err = db.Exec(helpers.InsertProduct, post.Product_id, post.Name, string(jsonStr), post.Sku, post.Category_id, price_round_off)

	if err != nil {
		helpers.CheckErr(err)
		helpers.SendErrResponse(helpers.Error, helpers.Query, w)
		helpers.LogError(err)
		return
	}

	response = typedefs.Json_Response{Type: helpers.Success, Message: helpers.InsertSuccess}

	json.NewEncoder(w).Encode(response)

}
