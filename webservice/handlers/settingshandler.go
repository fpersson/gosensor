package handlers

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

// SettingsHandler handles HTTP requests to display and manage application settings.
// It implements http.Handler and uses slog for logging.
type SettingsHandler struct {
}

// NewSettingsHandler creates and returns a new SettingsHandler instance.
//
// Returns:
//   - *SettingsHandler: a new SettingsHandler handler.
func NewSettingsHandler() *SettingsHandler {
	return &SettingsHandler{}
}

// ServeHTTP processes an incoming HTTP request and renders the settings page.
// It retrieves system information, application settings, and dynamically
// generates an HTML page using embedded templates.
//
// Parameters:
//   - w: http.ResponseWriter used to send the response.
//   - r: *http.Request received from the client.
//
// The rendered page includes settings, OS info, navigation, and footer.
func (settingsHandler *SettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idxPage := model.IndexPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	// Log error if unable to parse OS release information
	if err != nil {
		slog.Info(err.Error())
	}

	// Populate footer data with OS information
	idxPage.FooterData.OsString = osinfo["NAME"]
	idxPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	// Generate navigation menu based on the current URL path
	idxPage.NavPages = GetMenu(r.URL.Path)

	// Fetch all settings and log error if any
	data, err := model.ListAllSettings()
	if err != nil {
		slog.Info(err.Error())
	}

	idxPage.Settings = data
	slog.Info("(OPEN): " + "templates/settings.html")

	// Define template files
	navbar := "templates/navbar.html"
	footer := "templates/footer.html"
	t, err := template.ParseFS(content, "templates/settings.html", navbar, footer)

	// Log error if template parsing fails
	if err != nil {
		slog.Info(err.Error())
	}

	// Execute the template and handle potential errors
	if err := t.Execute(w, idxPage); err != nil {
		slog.Info("Template execution failed: " + err.Error())
	}
}
