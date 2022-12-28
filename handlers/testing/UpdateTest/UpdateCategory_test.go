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
		"category_id": 1,
		"name":        "Foods",
	}

	CheckUpdateCategory(category, "{\"type\":\"success\",\"message\":\"database has been updated successfully!\"}\n", t)

	category = map[string]any{
		"category_id": 10,
		"quantity":    "foods",
	}

	CheckUpdateCategory(category, "{\"type\":\"missing\",\"message\":\"Category id doesn't exist\"}\n", t)
}

func CheckUpdateCategory(category map[string]any, response string, t *testing.T) {
	json_product, err := json.Marshal(category)
	helpers.CheckErr(err)

	request_body := bytes.NewBuffer(json_product)
	req, err := http.NewRequest("PUT", "http://localhost:7171/updatecategory/", request_body)
	helpers.CheckErr(err)

	res, err := http.DefaultClient.Do(req)
	helpers.CheckErr(err)

	bodyBytes, err := io.ReadAll(res.Body)

	if string(bodyBytes) != response {
		t.Errorf("Expected: %v, Got: %v", response, string(bodyBytes))
	}
}
