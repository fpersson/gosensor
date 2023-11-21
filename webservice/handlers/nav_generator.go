package handlers

import "github.com/fpersson/gosensor/webservice/model"

func GetMenu(current_url string) (result model.NavPages) {
	tempnav := model.NavPages{}

	var content = model.NavPage{}

	content.Name = "Status"
	content.Url = "/sensor/status.html"
	content.IsActive = false
	tempnav.NavPage = append(tempnav.NavPage, content)

	content.Name = "Settings"
	content.Url = "/sensor/settings.html"
	content.IsActive = false
	tempnav.NavPage = append(tempnav.NavPage, content)

	content.Name = "Log"
	content.Url = "/sensor/log.html"
	content.IsActive = false
	tempnav.NavPage = append(tempnav.NavPage, content)

	content.Name = "Restart"
	content.Url = "/sensor/restart.html"
	content.IsActive = false
	tempnav.NavPage = append(tempnav.NavPage, content)

	for i, item := range tempnav.NavPage {
		if item.Url == current_url {
			tempnav.NavPage[i].IsActive = true
		}
	}

	return tempnav
}
