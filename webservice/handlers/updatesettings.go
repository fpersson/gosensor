package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fpersson/gosensor/webservice/model"
	"golang.org/x/exp/slog"
)

// UpdateSettings Handeler
type UpdateSettings struct {
	log *slog.Logger
}

// UpdateSettings function returning reference to UpdateSettings handle
func NewUpdateSettings(log *slog.Logger) *UpdateSettings {
	return &UpdateSettings{log}
}

func (updateSettings *UpdateSettings) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	updateSettings.log.Info("Update settings.")
	settings := &model.Settings{}

	if r.Method == "GET" {
		updateSettings.log.Info("TODO: implement GET for UpdateSettings")
		http.Redirect(w, r, "/", 301)
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
			updateSettings.log.Info(err.Error())
		}

		err = ioutil.WriteFile(model.SettingsPath, data, 0666)

		if err != nil {
			updateSettings.log.Info(err.Error())
		}

		http.Redirect(w, r, "/index.html", 301)
	}
}
