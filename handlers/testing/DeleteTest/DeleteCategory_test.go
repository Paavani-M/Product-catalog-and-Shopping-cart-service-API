package handlers

import (
	"io"
	"net/http"
	"testing"
)

func TestDeleteCategoryNotExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:7171/deletecategory/22/", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	// Check the response body, if necessary
	// ...

	expected := "{\"type\":\"missing\",\"message\":\"Category id doesn't exist\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}

func TestDeleteCategoryExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:7171/deletecategory/1/", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	// Check the response body, if necessary
	// ...

	expected := "{\"type\":\"success\",\"message\":\"Category type has been deleted successfully!\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}
