package model

var HttpDir string

// FooterData struct for info to show in footer
type FooterData struct {
	OsString  string
	OsVersion string
}

type NavPage struct {
	Name     string
	Url      string
	IsActive bool
}

type NavPages struct {
	NavPage []NavPage
}

type Settings struct {
	Sensor string `json:"sensor"`
	Name   string `json:"name"`
	//Influx  Influx_conf  `json:"influx"`
	///Grafana Grafana_conf `json:"grafana"`
}

// StatusPage datastruct for the status page
type StatusPage struct {
	NavPages   NavPages   //This is needed on every page
	FooterData FooterData //This is needed on every page
	//SystemdStatus SystemdStatus
}

// IndexPage Data struct for index page used for settings page
type IndexPage struct {
	Settings   Settings
	FooterData FooterData //This is needed on every page
	NavPages   NavPages   //This is needed on every page
}
