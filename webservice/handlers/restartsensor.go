package handlers

import (
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
)

// RestartSensor handles HTTP requests to restart the sensor service.
// It uses a logger to log information and errors during request processing.
type RestartSensor struct {
	log *slog.Logger
}

// NewRestartSensor creates and returns a new RestartSensor instance.
// It takes a logger as a parameter to enable logging.
func NewRestartSensor(log *slog.Logger) *RestartSensor {
	return &RestartSensor{log}
}

// ServeHTTP processes an incoming HTTP request to restart the sensor service.
// If the request method is GET, it redirects to the index page. For other methods,
// it attempts to restart the sensor service and redirects to the index page upon success.
//
// Parameters:
//   - w: The HTTP response writer used to send the response.
//   - r: The HTTP request received from the client.
func (restartSensor *RestartSensor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	restartSensor.log.Info("Restart sensor.")
	if r.Method == "GET" {
		restartSensor.log.Info("TODO: implement GET for RestartSensor")
		http.Redirect(w, r, "/index.html", http.StatusMovedPermanently)
	} else {
		restartSensor.log.Info("Restarting sensor.")
		err := syscmd.CmdRestart(syscmd.Tempsensorservice)
		if err != nil {
			restartSensor.log.Error("Failed to restart sensor service", "error", err)
			http.Error(w, "Failed to restart sensor service", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index.html", http.StatusMovedPermanently)
	}
}
