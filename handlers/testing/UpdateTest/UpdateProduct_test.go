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
		"product_id": 9,
		"name":       "Mac Book",
		"specification": map[string]any{
			"color":           "black",
			"special feature": "long battery life",
		},
		"sku":   "M1a78B",
		"price": 799999,
	}

	CheckUpdateProduct(product_master, "{\"type\":\"Success\",\"message\":\"Database has been updated successfully!\"}\n", t)

	//updating only one field
	product_master = map[string]any{"product_id": 9, "name": "MacBookpro", "price": 109000}
	CheckUpdateProduct(product_master, "{\"type\":\"Success\",\"message\":\"Database has been updated successfully!\"}\n", t)

	//product id not exists
	product_master = map[string]any{"product_id": 1277, "name": "Jacket", "price": 1099}
	CheckUpdateProduct(product_master, "{\"type\":\"Missing\",\"message\":\"Id doesn't exists!\"}\n", t)

}

func CheckUpdateProduct(product_master map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(product_master)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:7172/updateproduct/", request_body)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if string(bodyBytes) != response {
		t.Errorf("Expected: %v, Got: %v", response, string(bodyBytes))
	}
}
