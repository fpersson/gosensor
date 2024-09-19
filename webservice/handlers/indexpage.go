package handlers

import (
	"embed"
	"html/template"
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

type IndexPage struct {
	log *slog.Logger
}

func NewIndexPage(log *slog.Logger) *IndexPage {
	return &IndexPage{log}
}

//go:embed templates
var content embed.FS

func (indexPage *IndexPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	indexPage.log.Info("Open: " +  "templates/statusPage.html")
	statusPage := model.StatusPage{}

	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		indexPage.log.Info(err.Error())
	}
	
	statusPage.FooterData.OsString = osinfo["NAME"]
	statusPage.FooterData.OsVersion = osinfo["VERSION_ID"]	
	
	statusPage.NavPages = GetMenu(r.URL.Path)
	indexPage.log.Info("Reading status",
		slog.String("url", r.URL.Path),
	)

	data, err := syscmd.ReadStatus()
	if err != nil {
		indexPage.log.Info(err.Error())
	}

	statusPage.SystemdStatus = *data
	var temp =  "templates/error.html"

	if data.Active {
		indexPage.log.Info("Systemd service works")
		temp =  "templates/message.html"
	}

	navbar :=  "templates/navbar.html"
	footer :=  "templates/footer.html"

	t, err := template.ParseFS(content, "templates/statusPage.html", temp, navbar, footer)

	if err != nil {
		indexPage.log.Info(err.Error())
	}

	exec_err := t.Execute(w, statusPage)

	if exec_err != nil {
		indexPage.log.Info(exec_err.Error())
	}

}
