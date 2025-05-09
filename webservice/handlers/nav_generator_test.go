package handlers

import (
	"testing"

	"github.com/fpersson/gosensor/webservice/model"
)

// TestGetMenuStatusActive tests if the status page is active and all other pages are inactive.
func TestGetMenuStatusActive(t *testing.T) {
	// Helper function to create a NavPages structure
	createNavPages := func(activePage string) model.NavPages {
		result := model.NavPages{}
		pages := []struct {
			Name     string
			Url      string
			IsActive bool
		}{
			{"Status", "/status.html", activePage == "/status.html"},
			{"Settings", "/settings.html", activePage == "/settings.html"},
			{"Log", "/log.html", activePage == "/log.html"},
			{"Restart", "/restart.html", activePage == "/restart.html"},
		}

		for _, page := range pages {
			result.NavPage = append(result.NavPage, model.NavPage{
				Name:     page.Name,
				Url:      page.Url,
				IsActive: page.IsActive,
			})
		}
		return result
	}

	// Define test cases
	cases := []struct {
		in    string
		wants model.NavPages
	}{
		{"/status.html", createNavPages("/status.html")},
	}

	// Execute test cases
	for _, c := range cases {
		got := GetMenu(c.in)
		for i, page := range got.NavPage {
			if page.IsActive != c.wants.NavPage[i].IsActive {
				t.Errorf("GetMenu(%q) == %t for page %q, wants %t",
					c.in, page.IsActive, page.Name, c.wants.NavPage[i].IsActive)
			}
		}
	}
}
