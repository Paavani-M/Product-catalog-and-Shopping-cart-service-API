package handlers

import (
	"encoding/json"
	"net/http"

	"task.com/handlers/query_helpers"
	"task.com/helpers"
)

func AddItemstoCart(w http.ResponseWriter, r *http.Request) {

	response := []map[string]any{}
	request_body := []map[string]int{}

	reference_id := r.URL.Query().Get("ref")

	err := json.NewDecoder(r.Body).Decode(&request_body)
	if err != nil {
		helpers.LogError(err)
		return
	}

	for _, v := range request_body {
		new_response_item := map[string]any{}
		product_id := v["product_id"]
		quantity := v["quantity"]

		if product_id <= 0 || quantity <= 0 {
			helpers.SendErrResponse(helpers.Error, helpers.ValidInput, w)
			return
		}

		new_response_item["product_id"] = product_id
		new_response_item["quantity"] = quantity
		res := query_helpers.AddItemtoCart(reference_id, product_id, quantity)
		new_response_item["message"] = res.Message
		new_response_item["type"] = res.Type

		response = append(response, new_response_item)
	}

	json.NewEncoder(w).Encode(response)
}
