package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

func Get(t *testing.T, path string) *http.Response {
	resp, err := http.Get(baseURL + path)
	if err != nil {
		t.Fatalf("GET request to %s failed: %v", path, err)
	}
	return resp
}

func Post(t *testing.T, path string, body interface{}) *http.Response {
	data, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal JSON for POST request to %s: %v", path, err)
	}

	resp, err := http.Post(baseURL+path, "application/json", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("POST request to %s failed: %v", path, err)
	}
	return resp
}

func Put(t *testing.T, path string, body interface{}) *http.Response {
	data, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to marshal JSON for PUT request to %s: %v", path, err)
	}

	req, err := http.NewRequest(http.MethodPut, baseURL+path, bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Failed to create PUT request to %s: %v", path, err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("PUT request to %s failed: %v", path, err)
	}
	return resp
}
