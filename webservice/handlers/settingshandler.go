package handlers

import (
	"html/template"
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

// SettingsHandler handles HTTP requests to display and manage application settings.
// It uses a logger to log information and errors during request processing.
type SettingsHandler struct {
	log *slog.Logger
}

// NewSettingsHandler creates and returns a new SettingsHandler instance.
// It takes a logger as a parameter to enable logging.
func NewSettingsHandler(log *slog.Logger) *SettingsHandler {
	return &SettingsHandler{log}
}

// ServeHTTP processes an incoming HTTP request and renders the settings page.
// It retrieves system information, application settings, and dynamically
// generates an HTML page using embedded templates.
//
// Parameters:
//   - w: The HTTP response writer used to send the response.
//   - r: The HTTP request received from the client.
func (settingsHandler *SettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idxPage := model.IndexPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	// Log error if unable to parse OS release information
	if err != nil {
		settingsHandler.log.Info(err.Error())
	}

	// Populate footer data with OS information
	idxPage.FooterData.OsString = osinfo["NAME"]
	idxPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	// Generate navigation menu based on the current URL path
	idxPage.NavPages = GetMenu(r.URL.Path)

	// Fetch all settings and log error if any
	data, err := model.ListAllSettings()
	if err != nil {
		settingsHandler.log.Info(err.Error())
	}

	idxPage.Settings = data
	settingsHandler.log.Info("(OPEN): " + "templates/settings.html")

	// Define template files
	navbar := "templates/navbar.html"
	footer := "templates/footer.html"
	t, err := template.ParseFS(content, "templates/settings.html", navbar, footer)

	// Log error if template parsing fails
	if err != nil {
		settingsHandler.log.Info(err.Error())
	}

	// Execute the template and handle potential errors
	if err := t.Execute(w, idxPage); err != nil {
		settingsHandler.log.Info("Template execution failed: " + err.Error())
	}
}
