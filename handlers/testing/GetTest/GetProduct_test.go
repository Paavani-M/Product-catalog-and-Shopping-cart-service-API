package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"task.com/helpers"
)

func TestGetProductExists(t *testing.T) {

	resp, err := http.Get("http://localhost:7172/getproduct/1/")
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

func TestGetProductNotExists(t *testing.T) {

	resp, err := http.Get("http://localhost:7172/getproduct/222/")
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

	if response != "Id doesn't exists!" {
		t.Errorf("Expected Id doesn't exists!, got %s", response)
	}

}
