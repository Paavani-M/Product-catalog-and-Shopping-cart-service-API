package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"task.com/helpers"
)

func TestUpdateCategory(t *testing.T) {
	category := map[string]any{
		"category_id": 2,
		"name":        "Cleaning things",
	}

	CheckUpdateCategory(category, "{\"type\":\"Success\",\"message\":\"Database has been updated successfully!\"}\n", t)

	category = map[string]any{
		"category_id": 1000,
		"quantity":    "foods",
	}

	CheckUpdateCategory(category, "{\"type\":\"Missing\",\"message\":\"Id doesn't exists!\"}\n", t)
}

func CheckUpdateCategory(category map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(category)
	if err != nil {
		helpers.CheckErr(err)
		helpers.LogError(err)
	}

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:7172/updatecategory/", request_body)
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
