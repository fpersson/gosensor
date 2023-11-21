package handlers

import (
	"testing"

	"github.com/fpersson/gosensor/webservice/model"
)

// TestGetMenuStatusActive test if status page is active and all other pages are inactive.
func TestGetMenuStatusActive(t *testing.T) {
	result := model.NavPages{}
	var content = model.NavPage{}

	content.Name = "Status"
	content.Url = "/sensor/status.html"
	content.IsActive = true
	result.NavPage = append(result.NavPage, content)

	content.Name = "Settings"
	content.Url = "/sensor/settings.html"
	content.IsActive = false
	result.NavPage = append(result.NavPage, content)

	content.Name = "Log"
	content.Url = "/sensor/log.html"
	content.IsActive = false
	result.NavPage = append(result.NavPage, content)

	content.Name = "Restart"
	content.Url = "/sensor/restart.html"
	content.IsActive = false
	result.NavPage = append(result.NavPage, content)

	cases := []struct {
		in    string
		wants model.NavPages
	}{
		{"/sensor/status.html", result},
	}

	for _, c := range cases {
		got := GetMenu(c.in)
		if got.NavPage[0].IsActive != c.wants.NavPage[0].IsActive {
			t.Errorf("GetMenu(%q) == %t, wants %t", c.in, got.NavPage[0].IsActive, c.wants.NavPage[0].IsActive)
		}
	}

}
