package handlers

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetCategory(t *testing.T) {
	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Get("http://localhost:7171/getcategory/")
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}
