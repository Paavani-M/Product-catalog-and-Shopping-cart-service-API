package handlers

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestDeleteCartNotExists(t *testing.T) {

	data := []byte(`{"product_id":1, "reference_id":"a9u7hb"}`)

	req, err := http.NewRequest("DELETE", "http://localhost:7171/deletecart/", bytes.NewBuffer(data))
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

	expected := "{\"type\":\"missing\",\"message\":\"product id or reference_id doesn't exists\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}

func TestDeleteCartExists(t *testing.T) {

	data := []byte(`{"product_id":111, "reference_id":"1f45bb50-3f65-423d-b9c9-8daf85b29e3b"}`)

	req, err := http.NewRequest("DELETE", "http://localhost:7171/deletecart/", bytes.NewBuffer(data))
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

	expected := "{\"type\":\"success\",\"message\":\"Deleted successfully!\"}\n"

	bodyBytes, err := io.ReadAll(resp.Body)

	if string(bodyBytes) != expected {
		t.Errorf("unexpected: got %s, want %s", string(bodyBytes), expected)
	}

}
