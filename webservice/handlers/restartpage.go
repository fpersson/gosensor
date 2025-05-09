package handlers

import (
	"html/template"
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

// Reboot handles HTTP requests to display the reboot page.
// It implements http.Handler and uses slog for logging.
type Reboot struct {
}

// NewReboot creates and returns a new Reboot instance.
//
// Returns:
//   - *Reboot: a new Reboot handler.
func NewReboot() *Reboot {
	return &Reboot{}
}

// ServeHTTP processes an incoming HTTP request and renders the reboot page.
// It retrieves system information, generates navigation menus, and dynamically
// generates an HTML page using embedded templates.
//
// Parameters:
//   - w: http.ResponseWriter used to send the response.
//   - r: *http.Request received from the client.
//
// The rendered page includes reboot info, OS info, navigation, and footer.
func (reboot *Reboot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Info("(OPEN): " + "templates/restartPage.html")
	rebootPage := model.RebootPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	// Log and handle error if OS information cannot be parsed
	if err != nil {
		slog.Error("Failed to parse OS release info", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rebootPage.FooterData.OsString = osinfo["NAME"]
	rebootPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	rebootPage.NavPages = GetMenu(r.URL.Path)

	navbar := "templates/navbar.html"
	footer := "templates/footer.html"
	t, err := template.ParseFS(content, "templates/restartPage.html", navbar, footer)

	// Log and handle error if templates cannot be parsed
	if err != nil {
		slog.Error("Failed to parse templates", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template and write the response
	err = t.Execute(w, rebootPage)
	if err != nil {
		slog.Error("Failed to execute template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
