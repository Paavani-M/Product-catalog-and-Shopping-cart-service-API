package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"task.com/helpers"
	"task.com/typedefs"
)

func TestInsertCart(t *testing.T) {

	resp, err := http.Post("http://localhost:7172/addtocart?ref=585f2ef9-ddc1-435e-9d69-0ed12dc9ae29&product_id=6&quantity=1", "application/json", nil)
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

	if response.Type != "Success" || response.Message != "Added to Cart!" {
		t.Errorf("Expected response body 'Success Added to Cart!', got '%s %s'", response.Type, response.Message)
	}

}

// reference_id missing
func TestInsertRefNotValid(t *testing.T) {

	resp, err := http.Post("http://localhost:7172/addtocart?product_id=111&quantity=2", "application/json", nil)
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

	if response.Type != "Missing" || response.Message != "Reference ID not passed" {
		t.Errorf("Expected response body 'Reference ID not passed', got '%s %s'", response.Type, response.Message)
	}

}

// Insufficient or no stock
func TestInsertNotCart(t *testing.T) {

	resp, err := http.Post("http://localhost:7172/addtocart?ref=1f45bb50-3f65-423d-b9c9-8daf85b29e3b&product_id=2&quantity=4", "application/json", nil)
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

	if response.Type != "Insufficient" {
		t.Errorf("Expected response body 'Insufficient or No stock', got '%s %s'", response.Type, response.Message)
	}

}
