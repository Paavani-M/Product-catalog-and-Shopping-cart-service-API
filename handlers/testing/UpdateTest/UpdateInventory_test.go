package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"task.com/helpers"
)

func TestUpdateInventory(t *testing.T) {
	inventory := map[string]any{
		"product_id": 112,
		"quantity":   9,
	}

	CheckUpdateInventory(inventory, "{\"type\":\"success\",\"message\":\"Database has been updated successfully!\"}\n", t)

	inventory = map[string]any{
		"product_id": 1,
		"quantity":   10,
	}

	CheckUpdateInventory(inventory, "{\"type\":\"missing\",\"message\":\"Inventory id doesn't exist\"}\n", t)
}

func CheckUpdateInventory(inventory map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(inventory)
	helpers.CheckErr(err)

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:7171/updateinventory/", request_body)
	helpers.CheckErr(err)

	res, err := http.DefaultClient.Do(req)
	helpers.CheckErr(err)

	bodyBytes, err := io.ReadAll(res.Body)

	if string(bodyBytes) != response {
		t.Errorf("Expected: %v, Got: %v", response, string(bodyBytes))
	}
}
