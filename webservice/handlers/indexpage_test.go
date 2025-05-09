package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock syscmd and model if needed for more advanced tests.

func TestIndexPage_ServeHTTP(t *testing.T) {
	handler := NewIndexPage()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	// Check that HTML is returned (simple check)
	contentType := rr.Header().Get("Content-Type")
	if contentType != "" && contentType != "text/html; charset=utf-8" {
		t.Errorf("unexpected Content-Type: %s", contentType)
	}
}
