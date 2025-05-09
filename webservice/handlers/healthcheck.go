package handlers

import (
	"encoding/json"
	"net/http"

	"log/slog"
)

// Result represents the status of the health check response.
// The "result" field in the JSON output indicates the health status of the service.
type Result struct {
	Status string `json:"result"`
}

// HealthCheck implements http.Handler for the health check endpoint.
// It logs health check requests using slog.
type HealthCheck struct {
}

// NewHealthCheck creates a new HealthCheck handler.
// Logs creation of the handler using slog.
//
// Returns:
//   - *HealthCheck: a new HealthCheck handler.
func NewHealthCheck() *HealthCheck {
	slog.Info("New HealthCheck created")
	return &HealthCheck{}
}

// ServeHTTP handles HTTP GET requests for the health check endpoint.
// Responds with a JSON object containing the health status of the service.
//
// Parameters:
//   - w: http.ResponseWriter to write the response.
//   - r: *http.Request representing the incoming request.
func (healthCheck *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Info("HealthCheck called")
	result := &Result{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
