package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"log/slog"

	"github.com/fpersson/gosensor/webservice/model"
)

// UpdateSettings handles HTTP requests to update application settings.
// It implements http.Handler and uses slog for logging.
type UpdateSettings struct {
}

// NewUpdateSettings creates and returns a new UpdateSettings instance.
//
// Returns:
//   - *UpdateSettings: a new UpdateSettings handler.
func NewUpdateSettings() *UpdateSettings {
	return &UpdateSettings{}
}

// ServeHTTP processes an incoming HTTP request to update application settings.
// If the request method is GET, it redirects to the home page. For other methods,
// it updates the settings based on the form values and writes them to a file.
//
// Parameters:
//   - w: http.ResponseWriter used to send the response.
//   - r: *http.Request received from the client.
//
// On error, responds with HTTP 500 and an error message.
func (updateSettings *UpdateSettings) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Info("Update settings.")
	settings := &model.Settings{}

	if r.Method == "GET" {
		slog.Info("TODO: implement GET for UpdateSettings")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		settings.Sensor = r.FormValue("sensor")
		settings.Name = r.FormValue("name")
		settings.Influx.Host = r.FormValue("host")
		settings.Influx.Token = r.FormValue("token")
		settings.Influx.Apiorg = r.FormValue("api")
		settings.Influx.Bucket = r.FormValue("bucket")
		settings.Influx.Interval = r.FormValue("interval")
		settings.Grafana.Host = r.FormValue("grafana-host")

		data, err := json.MarshalIndent(&settings, "", "	")
		if err != nil {
			slog.Info("Failed to marshal settings: " + err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = os.WriteFile(model.SettingsPath, data, 0666)
		if err != nil {
			slog.Info("Failed to write settings file: " + err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index.html", http.StatusMovedPermanently)
	}
}
