// Package libsensor provides utilities for reading sensor data and posting it
// to an InfluxDB instance. It includes functions for sensor data parsing and
// database interaction.
package libsensor

import (
	"log/slog"
	"os"
	"time"

	libsettings "github.com/fpersson/gosensor/libsettings"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// failcount tracks the number of consecutive write errors to InfluxDB.
var failcount int = 0

// Post sends a data point to an InfluxDB instance.
//
// Parameters:
//   - dbconf: The settings containing InfluxDB configuration.
//   - data: The data value to be posted (e.g., temperature).
//
// Returns:
//   - error: An error object if any issues occur during the operation.
func Post(dbconf libsettings.Settings, data float64) error {
	client := influxdb2.NewClient(dbconf.Influx.Host, dbconf.Influx.Token)
	defer client.Close()

	writeAPI := client.WriteAPI(dbconf.Influx.Apiorg, dbconf.Influx.Bucket)

	point := influxdb2.NewPoint(dbconf.Name,
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"last": data},
		time.Now())

	errorsCh := writeAPI.Errors()
	go handleWriteErrors(errorsCh)

	writeAPI.WritePoint(point)
	writeAPI.Flush()

	return nil
}

// handleWriteErrors processes errors from the InfluxDB write API.
func handleWriteErrors(errorsCh <-chan error) {
	for err := range errorsCh {
		failcount++
		slog.Error("write error", slog.Int("error count", failcount), slog.String("error", err.Error()))
		if failcount > 10 {
			slog.Error("too many write errors, exiting")
			os.Exit(1)
		}
	}
}
