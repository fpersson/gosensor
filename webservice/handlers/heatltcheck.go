package handlers

import (
	"encoding/json"
	"net/http"

	"log/slog"
)

// Result represents the status of the health check response.
// It contains a single field "Status" which indicates the health status of the service.
type Result struct {
	Status string `json:"result"`
}

// HealthCheck handles requests for service health.
// It implements the http.Handler interface and logs health check requests.
type HealthCheck struct {
	logger *slog.Logger
}

// NewHealthCheck creates a new instance of HealthCheck.
// It takes a logger as a parameter to log health check events.
func NewHealthCheck(logger *slog.Logger) *HealthCheck {
	logger.Info("New HealthCheck created")
	return &HealthCheck{logger}
}

// ServeHTTP handles HTTP requests for the health check endpoint.
// It responds with a JSON object containing the health status of the service.
func (healthCheck *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthCheck.logger.Info("HealthCheck called")
	result := &Result{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
