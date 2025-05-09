package handlers

import (
	"html/template"
	"net/http"

	"log/slog"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
)

// LogHandle handles HTTP requests to display log information.
// It uses a logger to log information and errors during request processing.
type LogHandle struct {
	log *slog.Logger
}

// NewLogHandle creates and returns a new LogHandle instance.
// It takes a logger as a parameter to enable logging.
func NewLogHandle(log *slog.Logger) *LogHandle {
	return &LogHandle{log}
}

// ServeHTTP processes an incoming HTTP request and renders a log page.
// It retrieves system logs, reads log files or systemd logs, and dynamically
// generates an HTML page using embedded templates.
//
// Parameters:
//   - w: The HTTP response writer used to send the response.
//   - r: The HTTP request received from the client.
func (logHandle *LogHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data = &model.AllMessages{}
	logHandle.log.Info("(OPEN): " + "templates/logPage.html")

	// Initialize the log page structure
	logPage := model.LogPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	// Log and handle error if OS information cannot be parsed
	if err != nil {
		logHandle.log.Error("Failed to parse OS release info", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Populate footer data with OS information
	logPage.FooterData.OsString = osinfo["NAME"]
	logPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	// Generate navigation menu based on the current URL path
	logPage.NavPages = GetMenu(r.URL.Path)

	logHandle.log.Info("Reading logs")
	if model.LogDir == "" {
		// Read logs from systemd if no log directory is specified
		logHandle.log.Info("Reading from systemd")
		logContent, err := syscmd.ReadLog()
		if err != nil {
			logHandle.log.Error("Failed to read logs from systemd", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data = logContent
	} else {
		// Read logs from the specified log directory
		logHandle.log.Info("Read log:", "file", model.LogDir)
		logContent, err := syscmd.ReadLogFile(model.LogDir)
		if err != nil {
			logHandle.log.Error("Failed to read logs from file", "file", model.LogDir, "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		data = logContent
	}

	// Assign the retrieved log data to the log page structure
	logPage.AllMessages = *data

	// Define template file paths
	navbar := "templates/navbar.html"
	footer := "templates/footer.html"

	// Parse the templates
	t, err := template.ParseFS(content, "templates/logPage.html", navbar, footer)
	if err != nil {
		logHandle.log.Error("Failed to parse templates", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template and write the response
	if err := t.Execute(w, logPage); err != nil {
		logHandle.log.Error("Failed to execute template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
