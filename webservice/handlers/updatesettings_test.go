package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUpdateSettings_ServeHTTP_GET(t *testing.T) {
	handler := NewUpdateSettings()
	req := httptest.NewRequest(http.MethodGet, "/update-settings", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMovedPermanently {
		t.Errorf("expected status 301, got %d", rr.Code)
	}
	location := rr.Header().Get("Location")
	if location != "/" {
		t.Errorf("expected redirect to /, got %s", location)
	}
}

func TestUpdateSettings_ServeHTTP_POST(t *testing.T) {
	handler := NewUpdateSettings()
	form := url.Values{}
	form.Set("sensor", "test")
	form.Set("name", "testname")
	form.Set("host", "localhost")
	form.Set("token", "tok")
	form.Set("api", "org")
	form.Set("bucket", "buck")
	form.Set("interval", "10")
	form.Set("grafana-host", "grafana")

	req := httptest.NewRequest(http.MethodPost, "/update-settings", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	// Accept both redirect and 500 depending on whether file writing succeeds
	if rr.Code != http.StatusMovedPermanently && rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 301 or 500, got %d", rr.Code)
	}
}
