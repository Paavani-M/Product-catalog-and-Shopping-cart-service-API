package handlers

import (
	"io"
	"net/http"
	"testing"

	"task.com/helpers"
)

func TestDeleteCategoryNotExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:7172/deletecategory/222/", nil)
	if err != nil {
		t.Fatal(err)
		helpers.LogError(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		helpers.LogError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	expected := "{\"type\":\"Missing\",\"message\":\"Id doesn't exists!\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		helpers.LogError(err)
	}

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}

func TestDeleteCategoryExists(t *testing.T) {
	req, err := http.NewRequest("DELETE", "http://localhost:7172/deletecategory/8/", nil)
	if err != nil {
		t.Fatal(err)
		helpers.LogError(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		helpers.LogError(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	expected := "{\"type\":\"Success\",\"message\":\"Deleted successfully!\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		helpers.LogError(err)
	}

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}
