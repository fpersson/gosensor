package handlers

import (
	"net/http"

	"github.com/fpersson/gosensor/syscmd"
	"github.com/fpersson/gosensor/webservice/model"
	"golang.org/x/exp/slog"
)

type IndexPage struct {
	log *slog.Logger
}

func NewIndexPage(log *slog.Logger) *IndexPage {
	return &IndexPage{log}
}

func (indexPage *IndexPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	indexPage.log.Info("Open: " + model.HttpDir + "template/statusPage.html")
	statusPage := model.StatusPage{}

	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		indexPage.log.Info(err.Error())
	}

	statusPage.FooterData.OsString = osinfo["NAME"]
	statusPage.FooterData.OsVersion = osinfo["VERSION_ID"]

}
