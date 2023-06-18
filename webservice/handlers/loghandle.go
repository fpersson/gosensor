package handlers

import (
	"net/http"
	"text/template"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
	"golang.org/x/exp/slog"
)

type LogHandle struct {
	log *slog.Logger
}

func NewLogHandle(log *slog.Logger) *LogHandle {
	return &LogHandle{log}
}

func (logHandle *LogHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logHandle.log.Info("(OPEN): " + model.HttpDir + "templates/logPage.html")
	logPage := model.LogPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		logHandle.log.Info(err.Error())
	}

	logPage.FooterData.OsString = osinfo["NAME"]
	logPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	logPage.NavPages = GetMenu(r.URL.Path)

	logHandle.log.Info("Reading logs")
	data, err := syscmd.ReadLog()
	if err != nil {
		logHandle.log.Info(err.Error())
	}

	logPage.AllMessages = *data

	navbar := model.HttpDir + "templates/navbar.html"
	footer := model.HttpDir + "templates/footer.html"

	t, err := template.ParseFiles(model.HttpDir+"templates/logPage.html", navbar, footer)
	if err != nil {
		logHandle.log.Info(err.Error())
	}

	t.Execute(w, logPage)
}
