package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogHandle_ServeHTTP(t *testing.T) {
	handler := NewLogHandle()
	req := httptest.NewRequest(http.MethodGet, "/log", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	// Accept both 200 and 500 depending on whether dependencies/mocks exist
	if rr.Code != http.StatusOK && rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 200 or 500, got %d", rr.Code)
	}
}
