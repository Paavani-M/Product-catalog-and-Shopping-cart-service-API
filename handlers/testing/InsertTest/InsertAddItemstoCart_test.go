package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestInsertItemstoCart(t *testing.T) {
	data := []byte(`[{"product_id":111, "quantity":1},
	                {"product_id":116, "quantity":1}]`)

	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Post("http://localhost:7171/additemstocart?ref=585f2ef9-ddc1-435e-9d69-0ed12dc9ae29", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// response := typedefs.Json_Response{}
	// err = json.Unmarshal(body, &response)

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	//expected := "{\"type\":\"success\",\"message\":\"database has been updated successfully!\"}\n"

	expected := "[{\"message\":\"Added to cart!\",\"product_id\":113,\"quantity\":1,\"type\":\"success\"},{\"message\":\"Added to cart!\",\"product_id\":116,\"quantity\":1,\"type\":\"success\"}]\n"

	if string(body) != expected {
		t.Errorf("Expected %s, got %s", expected, string(body))
	}

}

// reference_id missing
// func TestInsertRefNotValid_(t *testing.T) {
// 	data := []byte(`{"product_id":111, "quantity":2}`)

// 	resp, err := http.Post("http://localhost:7171/additemstocart/", "application/json", bytes.NewBuffer(data))
// 	if err != nil {
// 		t.Errorf("Error making request: %v", err)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if err != nil {
// 		t.Errorf("Error reading response body: %v", err)
// 	}

// 	response := typedefs.Json_Response{}
// 	err = json.Unmarshal(body, &response)

// 	if resp.StatusCode != 200 {
// 		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
// 	}

// 	if response.Type != "missing" || response.Message != "reference id has not passed" {
// 		t.Errorf("Expected response body 'invalid Invalid reference id has been passed', got '%s %s'", response.Type, response.Message)
// 	}

// }

// Insufficient or no stock
func TestInsertItemsNotCart(t *testing.T) {
	data := []byte(`[{
		              "product_id":125,
	                  "quantity":4
					},
	                { 
					  "product_id":119,
					  "quantity":4
					  }]`)

	resp, err := http.Post("http://localhost:7171/additemstocart?ref=1f45bb50-3f65-423d-b9c9-8daf85b29e3b", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// response := typedefs.Json_Response{}
	// err = json.Unmarshal(body, &response)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	//expected := "[{\"message\":\"Added to cart!\",\"product_id\":123,\"quantity\":2,\"type\":\"Insufficient\"},{\"message\":\"Added to cart!\",\"product_id\":124,\"quantity\":2,\"type\":\"Insufficient\"}]\n"

	// expected := "[{\"message\":\"Available quantity:1, Enough Stock doesn't exists\",\"product_id\":119,\"quantity\":10,\"type\":\"Insuifficient\"},{\"message\":\"Available quantity:1, Enough Stock doesn't exists\",\"product_id\":125,\"quantity\":10,\"type\":\"Insuifficient\"}]\n"

	// expected := "[{\"product_id\":125,\"quantity\":4,\"type\":\"Insuifficient\"},{\"product_id\":112,\"quantity\":4,\"type\":\"Insuifficient\"}]\n"

	// if string(body) != expected {
	// 	t.Errorf("Error: Expected %s, got %s", expected, string(body))
	// }
}
