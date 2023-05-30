package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/exp/slog"
)

type Result struct {
	Status string `json:"result"`
}

type HealthCheck struct {
	logger *slog.Logger
}

func NewHealthCheck(logger *slog.Logger) *HealthCheck {
	logger.Info("HealthCheck cStor called")
	return &HealthCheck{logger}
}

func (healthCheck *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthCheck.logger.Info("HealthCheck called")
	result := &Result{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}
