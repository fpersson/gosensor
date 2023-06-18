package handlers

import (
	"net/http"
	"text/template"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
	"golang.org/x/exp/slog"
)

type SettingsHandler struct {
	log *slog.Logger
}

func NewSettingsHandler(log *slog.Logger) *SettingsHandler {
	return &SettingsHandler{log}
}

func (settingsHandler *SettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idxPage := model.IndexPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		settingsHandler.log.Info(err.Error())
	}

	idxPage.FooterData.OsString = osinfo["NAME"]
	idxPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	idxPage.NavPages = GetMenu(r.URL.Path)

	data, err := model.ListAllSettings()
	if err != nil {
		settingsHandler.log.Info(err.Error())
	}

	idxPage.Settings = data
	settingsHandler.log.Info("(OPEN): " + model.HttpDir + "templates/settings.html")

	navbar := model.HttpDir + "templates/navbar.html"
	footer := model.HttpDir + "templates/footer.html"
	t, err := template.ParseFiles(model.HttpDir+"templates/settings.html", navbar, footer)

	if err != nil {
		settingsHandler.log.Info(err.Error())
	}

	t.Execute(w, idxPage)
}
