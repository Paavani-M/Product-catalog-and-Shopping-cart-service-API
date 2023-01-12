package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"task.com/helpers"
)

func TestGetCartExists(t *testing.T) {

	resp, err := http.Get("http://localhost:7172/cart/get_cart/c07eca02-260d-4508-a3c2-976c877691b9/")
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

func TestGetCartNotExists(t *testing.T) {

	resp, err := http.Get("http://localhost:7172/cart/get_cart/222/")
	if err != nil {
		t.Errorf("Error making request: %v", err)
		helpers.LogError(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var response string

	err = json.Unmarshal(body, &response)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
		helpers.LogError(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response != "Reference ID doesn't exists!" {
		t.Errorf("Expected Reference ID doesn't exists!, got %s", response)
	}

}
