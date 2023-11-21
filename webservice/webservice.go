package webservice

import (
	"net/http"
	"time"

	"log/slog"

	"github.com/fpersson/gosensor/webservice/handlers"
)

type WebService struct {
	logger *slog.Logger
}

func NewWebService(logger *slog.Logger) *WebService {
	return &WebService{logger}
}

func (webservice *WebService) Start() {
	webservice.logger.Info("Start called.")

	serveMux := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8081",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	healtCheck := handlers.NewHealthCheck(webservice.logger)
	indexPage := handlers.NewIndexPage(webservice.logger)
	settingsPage := handlers.NewSettingsHandler(webservice.logger)
	updatePage := handlers.NewUpdateSettings(webservice.logger)
	logPage := handlers.NewLogHandle(webservice.logger)
	restartService := handlers.NewRestartSensor(webservice.logger)
	rebootPage := handlers.NewReboot(webservice.logger)
	serveMux.Handle("/healthcheck", healtCheck)
	serveMux.Handle("/health_check", healtCheck)
	serveMux.Handle("/", indexPage)
	serveMux.Handle("/index.html", indexPage)
	serveMux.Handle("/status.html", indexPage)
	serveMux.Handle("/settings.html", settingsPage)
	serveMux.Handle("/log.html", logPage)
	serveMux.Handle("/restart.html", rebootPage)
	serveMux.Handle("/update", updatePage)
	serveMux.Handle("/reboot_system", restartService)

	err := server.ListenAndServe()

	if err != nil {
		webservice.logger.Error(err.Error())
	}
}
