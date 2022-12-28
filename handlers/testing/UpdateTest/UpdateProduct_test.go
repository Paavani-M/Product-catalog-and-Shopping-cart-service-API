package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"task.com/helpers"
)

func TestUpdateProduct(t *testing.T) {
	product_master := map[string]any{
		"product_id": 119,
		"name":       "Boat wireless earphones",
		"specification": map[string]any{
			"color":           "blue",
			"special feature": "long battery life",
		},
		"sku":   "b1a78t",
		"price": 799,
	}

	CheckUpdateProduct(product_master, "{\"type\":\"success\",\"message\":\"database has been updated successfully!\"}\n", t)

	//updating only one field
	product_master = map[string]any{"product_id": 124, "name": "Jacket", "price": 1090}
	CheckUpdateProduct(product_master, "{\"type\":\"success\",\"message\":\"database has been updated successfully!\"}\n", t)

	//product id not exists
	product_master = map[string]any{"product_id": 127, "name": "Jacket", "price": 1099}
	CheckUpdateProduct(product_master, "{\"type\":\"missing\",\"message\":\"product id doesn't exists!\"}\n", t)

}

func CheckUpdateProduct(product_master map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(product_master)
	helpers.CheckErr(err)

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:7171/updateproduct/", request_body)
	helpers.CheckErr(err)

	res, err := http.DefaultClient.Do(req)
	helpers.CheckErr(err)

	bodyBytes, err := io.ReadAll(res.Body)

	if string(bodyBytes) != response {
		t.Errorf("Expected: %v, Got: %v", response, string(bodyBytes))
	}
}
