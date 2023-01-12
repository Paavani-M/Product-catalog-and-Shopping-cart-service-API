package handlers

import (
	"io/ioutil"
	"net/http"
	"testing"

	"task.com/helpers"
)

func TestGetCategory(t *testing.T) {

	resp, err := http.Get("http://localhost:7172/getcategory/")
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
