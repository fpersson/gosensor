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
	content.Url = "/status.html"
	content.IsActive = true
	result.NavPage = append(result.NavPage, content)

	content.Name = "Settings"
	content.Url = "/settings.html"
	content.IsActive = false
	result.NavPage = append(result.NavPage, content)

	content.Name = "Log"
	content.Url = "/log.html"
	content.IsActive = false
	result.NavPage = append(result.NavPage, content)

	content.Name = "Restart"
	content.Url = "/restart.html"
	content.IsActive = false
	result.NavPage = append(result.NavPage, content)

	cases := []struct {
		in    string
		wants model.NavPages
	}{
		{"/status.html", result},
	}

	for _, c := range cases {
		got := GetMenu(c.in)
		if got.NavPage[0].IsActive != c.wants.NavPage[0].IsActive {
			t.Errorf("GetMenu(%q) == %t, wants %t", c.in, got.NavPage[0].IsActive, c.wants.NavPage[0].IsActive)
		}
	}

}
