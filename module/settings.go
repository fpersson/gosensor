package module

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
