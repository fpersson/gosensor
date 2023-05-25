package handlers

import (
	"fmt"
	"net/http"

	"golang.org/x/exp/slog"
)

type HealthCheck struct {
	logger *slog.Logger
}

func NewHealthCheck(logger *slog.Logger) *HealthCheck {
	logger.Info("HealthCheck cStor called")
	return &HealthCheck{logger}
}

func (healthCheck *HealthCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthCheck.logger.Info("HealthCheck called")

	fmt.Fprintf(w, "OK")
}
