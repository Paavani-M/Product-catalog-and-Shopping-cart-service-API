package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"task.com/typedefs"
)

func TestInsertCart(t *testing.T) {

	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Post("http://localhost:7171/addtocart?ref=585f2ef9-ddc1-435e-9d69-0ed12dc9ae29&product_id=119&quantity=1", "application/json", nil)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.Json_Response{}
	err = json.Unmarshal(body, &response)

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Type != "success" || response.Message != "Added to cart!" {
		t.Errorf("Expected response body 'success added to cart!', got '%s %s'", response.Type, response.Message)
	}

}

// reference_id missing
func TestInsertRefNotValid(t *testing.T) {
	//data := []byte(`{"product_id":111, "quantity":2}`)

	resp, err := http.Post("http://localhost:7171/addtocart?product_id=111&quantity=2", "application/json", nil)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.Json_Response{}
	err = json.Unmarshal(body, &response)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Type != "missing" || response.Message != "reference id has not passed" {
		t.Errorf("Expected response body 'invalid Invalid reference id has been passed', got '%s %s'", response.Type, response.Message)
	}

}

// Insufficient or no stock
func TestInsertNotCart(t *testing.T) {

	resp, err := http.Post("http://localhost:7171/addtocart?ref=1f45bb50-3f65-423d-b9c9-8daf85b29e3b&product_id=124&quantity=4", "application/json", nil)
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.Json_Response{}
	err = json.Unmarshal(body, &response)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Type != "Insufficient" {
		t.Errorf("Expected response body 'Insufficient or No stock', got '%s %s'", response.Type, response.Message)
	}

}
