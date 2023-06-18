package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var SettingsPath string
var HttpDir string
var LogDir string

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
	Sensor  string       `json:"sensor"`
	Name    string       `json:"name"`
	Influx  Influx_conf  `json:"influx"`
	Grafana Grafana_conf `json:"grafana"`
}

type Influx_conf struct {
	Host     string `json:"host"`
	Token    string `json:"token"` //can be "username:password for influx v1"
	Apiorg   string `json:"api_org"`
	Bucket   string `json:"bucket"`
	Interval string `json:"interval"`
}

type Grafana_conf struct {
	Host string `json:"host"`
}

// StatusPage datastruct for the status page
type StatusPage struct {
	NavPages      NavPages   //This is needed on every page
	FooterData    FooterData //This is needed on every page
	SystemdStatus SystemdStatus
}

// IndexPage Data struct for index page used for settings page
type IndexPage struct {
	Settings   Settings
	FooterData FooterData //This is needed on every page
	NavPages   NavPages   //This is needed on every page
}

type AllMessages struct {
	LogMessages *[]string
}

type SystemdStatus struct {
	Active  bool
	Message *[]string
}

// LogPage Data struct for log page
type LogPage struct {
	FooterData  FooterData //This is needed on every page
	NavPages    NavPages   //This is needed on every page
	AllMessages AllMessages
}

type RebootPage struct {
	FooterData FooterData //This is needed on every page
	NavPages   NavPages   //This is needed on every page
}

// ListAllSettings @todo move this code and write a test
// TODO I know this part is duplicated but I keep it for now.
func ListAllSettings() (settings Settings, err error) {
	var result Settings

	jsonFile, err := os.Open(SettingsPath)

	if err != nil {
		return result, err
	}

	b, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(b, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}
