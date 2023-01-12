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

func TestInsertCategory(t *testing.T) {
	data := []byte(`{"category_id":13, "name":"Utensils"}`)

	resp, err := http.Post("http://localhost:7172/insertcategory/", "application/json", bytes.NewBuffer(data))
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
