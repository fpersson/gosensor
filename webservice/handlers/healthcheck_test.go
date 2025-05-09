package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck_ServeHTTP(t *testing.T) {
	handler := NewHealthCheck()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	var result Result
	err := json.NewDecoder(rr.Body).Decode(&result)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if result.Status != "OK" {
		t.Errorf("expected result 'OK', got '%s'", result.Status)
	}
}
