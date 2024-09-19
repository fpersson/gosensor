package handlers

import (
	"html/template"
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

type LogHandle struct {
	log *slog.Logger
}

func NewLogHandle(log *slog.Logger) *LogHandle {
	return &LogHandle{log}
}

func (logHandle *LogHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data = &model.AllMessages{}
	logHandle.log.Info("(OPEN): " + "templates/logPage.html")
	logPage := model.LogPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		logHandle.log.Info(err.Error())
	}

	logPage.FooterData.OsString = osinfo["NAME"]
	logPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	logPage.NavPages = GetMenu(r.URL.Path)

	logHandle.log.Info("Reading logs")
	if model.LogDir == "" {
		logHandle.log.Info("Reading from systemd")
		logContent, err := syscmd.ReadLog()
		if err != nil {
			logHandle.log.Info(err.Error())
		}
		data = logContent
	} else {
		logHandle.log.Info("Read log:", "file", model.LogDir)
		logContent, err := syscmd.ReadLogFile(model.LogDir)
		if err != nil {
			logHandle.log.Info(err.Error())
		}
		data = logContent
	}

	logPage.AllMessages = *data

	navbar := "templates/navbar.html"
	footer := "templates/footer.html"

	t, err := template.ParseFS(content, "templates/logPage.html", navbar, footer)
	if err != nil {
		logHandle.log.Info(err.Error())
	}

	t.Execute(w, logPage)
}
