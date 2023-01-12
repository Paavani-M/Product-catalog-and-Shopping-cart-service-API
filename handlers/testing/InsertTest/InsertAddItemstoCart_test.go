package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"task.com/helpers"
)

func TestInsertItemstoCart(t *testing.T) {
	data := []byte(`[{"product_id":1, "quantity":1},
	                {"product_id":3, "quantity":1}]`)

	resp, err := http.Post("http://localhost:7172/additemstocart?ref=585f2ef9-ddc1-435e-9d69-0ed12dc9ae29", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
		helpers.LogError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		helpers.LogError(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	expected := "[{\"message\":\"Added to Cart!\",\"product_id\":1,\"quantity\":1,\"type\":\"Success\"},{\"message\":\"Added to Cart!\",\"product_id\":3,\"quantity\":1,\"type\":\"Success\"}]\n"

	if string(body) != expected {
		t.Errorf("Expected %s, got %s", expected, string(body))
	}

}

// Insufficient or no stock
func TestInsertItemsNotCart(t *testing.T) {
	data := []byte(`[{
		              "product_id":4,
	                  "quantity":40
					},
	                { 
					  "product_id":5,
					  "quantity":40
					  }]`)

	resp, err := http.Post("http://localhost:7172/additemstocart?ref=1f45bb50-3f65-423d-b9c9-8daf85b29e3b", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
		helpers.LogError(err)
	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		helpers.LogError(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}
