package handlers

import (
	"embed"
	"html/template"
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

// IndexPage handles HTTP requests to display the service status.
// It uses a logger to log information and errors during request processing.
type IndexPage struct {
	log *slog.Logger
}

// NewIndexPage creates and returns a new IndexPage instance.
// It takes a logger as a parameter to enable logging.
func NewIndexPage(log *slog.Logger) *IndexPage {
	return &IndexPage{log}
}

//go:embed templates
var content embed.FS

// ServeHTTP processes an incoming HTTP request and renders a status page.
// It retrieves system information, reads the service status, and dynamically
// generates an HTML page using embedded templates.
//
// Parameters:
//   - w: The HTTP response writer used to send the response.
//   - r: The HTTP request received from the client.
func (indexPage *IndexPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	indexPage.log.Info("Open: " + "templates/statusPage.html")
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
	var temp = "templates/error.html"

	if data.Active {
		indexPage.log.Info("Systemd service works")
		temp = "templates/message.html"
	}

	navbar := "templates/navbar.html"
	footer := "templates/footer.html"

	t, err := template.ParseFS(content, "templates/statusPage.html", temp, navbar, footer)

	if err != nil {
		indexPage.log.Info(err.Error())
	}

	exec_err := t.Execute(w, statusPage)

	if exec_err != nil {
		indexPage.log.Info(exec_err.Error())
	}

}
