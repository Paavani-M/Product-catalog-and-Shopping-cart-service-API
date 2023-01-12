package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"task.com/helpers"
	"task.com/typedefs"
)

func TestInsertProduct(t *testing.T) {

	data := []byte(`{"product_id":1551,"name":"iour", "specification":{"origin":"usa"}, "sku":"i91u0", "category_id":3, "price":12}`)

	resp, err := http.Post("http://localhost:7172/insertproduct/", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
		helpers.LogError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		helpers.LogError(err)
	}

	response := typedefs.Json_Response{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		helpers.LogError(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Type != "Success" || response.Message != "Record has been inserted successfully!" {
		t.Errorf("Expected response body 'success Record has been inserted successfully!', got '%s %s'", response.Type, response.Message)
	}

}
