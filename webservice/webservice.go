package webservice

import (
	"net/http"
	"time"

	"github.com/fpersson/gosensor/webservice/handlers"
	"golang.org/x/exp/slog"
)

type WebService struct {
	logger *slog.Logger
}

func NewWebService(logger *slog.Logger) *WebService {
	return &WebService{logger}
}

func (webservice *WebService) Start() {
	webservice.logger.Info("Start called...")

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
	serveMux.Handle("/healthcheck", healtCheck)
	serveMux.Handle("/health_check", healtCheck)
	serveMux.Handle("/index.html", indexPage)

	err := server.ListenAndServe()

	if err != nil {
		webservice.logger.Error(err.Error())
	}
}
