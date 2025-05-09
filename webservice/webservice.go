package webservice

import (
	"net/http"
	"time"

	"log/slog"

	"github.com/fpersson/gosensor/webservice/handlers"
)

type WebService struct {
}

func NewWebService() *WebService {
	return &WebService{}
}

func (webservice *WebService) Start() {
	slog.Info("Start called.")

	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8081",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	healtCheck := handlers.NewHealthCheck()
	indexPage := handlers.NewIndexPage()
	settingsPage := handlers.NewSettingsHandler()
	updatePage := handlers.NewUpdateSettings()
	logPage := handlers.NewLogHandle()
	restartService := handlers.NewRestartSensor()
	rebootPage := handlers.NewReboot()
	serveMux.Handle("/healthcheck", healtCheck)
	serveMux.Handle("/health_check", healtCheck)
	serveMux.Handle("/index.html", indexPage)
	serveMux.Handle("/status.html", indexPage)
	serveMux.Handle("/settings.html", settingsPage)
	serveMux.Handle("/log.html", logPage)
	serveMux.Handle("/restart.html", rebootPage)
	serveMux.Handle("/update", updatePage)
	serveMux.Handle("/reboot_system", restartService)

	err := server.ListenAndServe()

	if err != nil {
		slog.Error(err.Error())
	}
}
