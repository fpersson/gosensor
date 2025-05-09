package handlers

import (
	"embed"
	"html/template"
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

// IndexPage handles HTTP requests to display the service status page.
// It implements http.Handler and uses slog for logging.
type IndexPage struct {
}

// NewIndexPage creates and returns a new IndexPage instance.
//
// Returns:
//   - *IndexPage: a new IndexPage handler.
func NewIndexPage() *IndexPage {
	return &IndexPage{}
}

//go:embed templates
var content embed.FS

// ServeHTTP processes an incoming HTTP request and renders a status page.
// It retrieves system information, reads the service status, and dynamically
// generates an HTML page using embedded templates.
//
// Parameters:
//   - w: http.ResponseWriter used to send the response.
//   - r: *http.Request received from the client.
//
// The rendered page includes system status, OS info, navigation, and footer.
func (indexPage *IndexPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Info("Open: " + "templates/statusPage.html")
	statusPage := model.StatusPage{}

	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		slog.Info(err.Error())
	}

	statusPage.FooterData.OsString = osinfo["NAME"]
	statusPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	statusPage.NavPages = GetMenu(r.URL.Path)
	slog.Info("Reading status",
		slog.String("url", r.URL.Path),
	)

	data, err := syscmd.ReadStatus()
	if err != nil {
		slog.Info(err.Error())
	}

	statusPage.SystemdStatus = *data
	var temp = "templates/error.html"

	if data.Active {
		slog.Info("Systemd service works")
		temp = "templates/message.html"
	}

	navbar := "templates/navbar.html"
	footer := "templates/footer.html"

	t, err := template.ParseFS(content, "templates/statusPage.html", temp, navbar, footer)

	if err != nil {
		slog.Info(err.Error())
	}

	exec_err := t.Execute(w, statusPage)

	if exec_err != nil {
		slog.Info(exec_err.Error())
	}

}
