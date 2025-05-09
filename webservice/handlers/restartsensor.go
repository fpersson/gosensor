package handlers

import (
	"log/slog"
	"net/http"

	"github.com/fpersson/gosensor/syscmd"
)

// RestartSensor handles HTTP requests to restart the sensor service.
// It implements http.Handler and uses slog for logging.
type RestartSensor struct {
}

// NewRestartSensor creates and returns a new RestartSensor instance.
//
// Returns:
//   - *RestartSensor: a new RestartSensor handler.
func NewRestartSensor() *RestartSensor {
	return &RestartSensor{}
}

// ServeHTTP processes an incoming HTTP request to restart the sensor service.
// If the request method is GET, it redirects to the index page. For other methods,
// it attempts to restart the sensor service and redirects to the index page upon success.
//
// Parameters:
//   - w: http.ResponseWriter used to send the response.
//   - r: *http.Request received from the client.
//
// On error, responds with HTTP 500 and an error message.
func (restartSensor *RestartSensor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Info("Restart sensor.")
	if r.Method == "GET" {
		slog.Info("TODO: implement GET for RestartSensor")
		http.Redirect(w, r, "/index.html", http.StatusMovedPermanently)
	} else {
		slog.Info("Restarting sensor.")
		err := syscmd.CmdRestart(syscmd.Tempsensorservice)
		if err != nil {
			slog.Error("Failed to restart sensor service", "error", err)
			http.Error(w, "Failed to restart sensor service", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index.html", http.StatusMovedPermanently)
	}
}
