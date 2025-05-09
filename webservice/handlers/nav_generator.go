package handlers

import "github.com/fpersson/gosensor/webservice/model"

// GetMenu generates a navigation menu with predefined pages and marks the
// current URL as active. It returns a `model.NavPages` structure containing
// the navigation data.
//
// Parameters:
//   - current_url: The URL of the current page, which will be marked as active.
//
// Returns:
//   - result: A `model.NavPages` structure containing the navigation menu.
func GetMenu(current_url string) (result model.NavPages) {
	// Helper function to create a navigation page
	createNavPage := func(name, url string, isActive bool) model.NavPage {
		return model.NavPage{
			Name:     name,
			Url:      url,
			IsActive: isActive,
		}
	}

	// Initialize navigation pages
	tempnav := model.NavPages{
		NavPage: []model.NavPage{
			createNavPage("Status", "/status.html", false),
			createNavPage("Settings", "/settings.html", false),
			createNavPage("Log", "/log.html", false),
			createNavPage("Restart", "/restart.html", false),
		},
	}

	// Mark the current URL as active
	for i, item := range tempnav.NavPage {
		if item.Url == current_url {
			tempnav.NavPage[i].IsActive = true
		}
	}

	return tempnav
}
