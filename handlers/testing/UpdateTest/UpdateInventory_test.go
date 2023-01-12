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
		"product_id": 10,
		"quantity":   7,
	}

	CheckUpdateInventory(inventory, "{\"type\":\"Success\",\"message\":\"Database has been updated successfully!\"}\n", t)

	inventory = map[string]any{
		"product_id": 1000,
		"quantity":   10,
	}

	CheckUpdateInventory(inventory, "{\"type\":\"Missing\",\"message\":\"Id doesn't exists!\"}\n", t)
}

func CheckUpdateInventory(inventory map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(inventory)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:7172/updateinventory/", request_body)
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
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	if string(bodyBytes) != response {
		t.Errorf("Expected: %v, Got: %v", response, string(bodyBytes))
	}
}
