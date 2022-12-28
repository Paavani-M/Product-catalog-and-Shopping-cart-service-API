package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"task.com/typedefs"
)

func TestInsertCategory(t *testing.T) {
	data := []byte(`{"category_id":7, "name":"essentials"}`)

	// Make a request to the API endpoint that triggers the insert function
	resp, err := http.Post("http://localhost:7171/insertcategory/", "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Error making request: %v", err)
	}

	// Read the response from the API
	body, err := ioutil.ReadAll(resp.Body)

	//fmt.Println("response", string(body))

	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	response := typedefs.Json_Response{}
	err = json.Unmarshal(body, &response)

	// Make assertions about the output of the function
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	if response.Type != "success" || response.Message != "Record has been inserted successfully!" {
		t.Errorf("Expected response body 'success Record has been inserted successfully!', got '%s %s'", response.Type, response.Message)
	}

}
