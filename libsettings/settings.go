// Package libsettings provides utilities for managing and parsing configuration settings.
// It includes functionality to read JSON configuration files and define the structure of settings.
package libsettings

// Settings represents the configuration for the sensor system.
// It includes sensor details, InfluxDB configuration, and Grafana configuration.
type Settings struct {
	Sensor  string       `json:"sensor"`  // The type of sensor being used.
	Name    string       `json:"name"`    // The name of the sensor.
	Influx  Influx_conf  `json:"influx"`  // Configuration for InfluxDB.
	Grafana Grafana_conf `json:"grafana"` // Configuration for Grafana.
}

// Influx_conf represents the configuration for connecting to an InfluxDB instance.
type Influx_conf struct {
	Host     string `json:"host"`     // The host address of the InfluxDB server.
	Token    string `json:"token"`    // Authentication token or "username:password" for InfluxDB v1.
	Apiorg   string `json:"api_org"`  // The organization name for InfluxDB API.
	Bucket   string `json:"bucket"`   // The bucket name in InfluxDB.
	Interval string `json:"interval"` // The data collection interval.
}

// Grafana_conf represents the configuration for connecting to a Grafana instance.
type Grafana_conf struct {
	Host string `json:"host"` // The host address of the Grafana server.
}
