package handlers

import (
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
)

type RestartSensor struct {
	log *slog.Logger
}

func NewRestartSensor(log *slog.Logger) *RestartSensor {
	return &RestartSensor{log}
}

func (restartSensor *RestartSensor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	restartSensor.log.Info("Restart sensor.")
	if r.Method == "GET" {
		restartSensor.log.Info("TODO: implement GET for RestartSensor")
		http.Redirect(w, r, "/index.html", 301)
	} else {
		restartSensor.log.Info("Restarting sensor.")
		err := syscmd.CmdRestart(syscmd.Tempsensorservice)
		if err != nil {
			restartSensor.log.Info(err.Error())
		}
		http.Redirect(w, r, "/index.html", 301)
	}
}
