package handlers

import (
	"html/template"
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

type Reboot struct {
	log *slog.Logger
}

func NewReboot(log *slog.Logger) *Reboot {
	return &Reboot{log}
}

func (reboot *Reboot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reboot.log.Info("(OPEN): " + model.HttpDir + "templates/restartPage.html")
	rebootPage := model.RebootPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		reboot.log.Info(err.Error())
	}

	rebootPage.FooterData.OsString = osinfo["NAME"]
	rebootPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	rebootPage.NavPages = GetMenu(r.URL.Path)

	navbar := model.HttpDir + "templates/navbar.html"
	footer := model.HttpDir + "templates/footer.html"
	t, err := template.ParseFiles(model.HttpDir+"templates/restartPage.html", navbar, footer)

	if err != nil {
		reboot.log.Info(err.Error())
	}

	t.Execute(w, rebootPage)
}
