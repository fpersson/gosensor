package libsensor

import (
	"log/slog"
	"os"
	"time"

	libsettings "github.com/fpersson/gosensor/libsettings"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var failcount int =0

func Post(dbconf libsettings.Settings, data float64) error {
	client := influxdb2.NewClient(dbconf.Influx.Host, dbconf.Influx.Token)
	writeAPI := client.WriteAPI(dbconf.Influx.Apiorg, dbconf.Influx.Bucket)

	point := influxdb2.NewPoint(dbconf.Name,
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"last": data},
		time.Now())

		errorsCh := writeAPI.Errors()
		go func() {
			for err := range errorsCh {
				failcount++
				slog.Error("write error", slog.Int("error count", failcount), slog.String("error", err.Error()))
				if failcount > 10 {
					os.Exit(1)
				}
			}
		}()

	writeAPI.WritePoint(point)
	writeAPI.Flush()
	
	client.Close()
	return nil
}
