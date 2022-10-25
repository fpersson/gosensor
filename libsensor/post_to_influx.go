package libsensor

import (
	"time"

	libsettings "github.com/fpersson/gosensor/libsettings"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func Post(dbconf libsettings.Settings, data float64) error {
	client := influxdb2.NewClient(dbconf.Influx.Host, dbconf.Influx.Token)
	writeAPI := client.WriteAPI(dbconf.Influx.Apiorg, dbconf.Influx.Bucket)

	point := influxdb2.NewPoint("sensor_1",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"last": data},
		time.Now())

	writeAPI.WritePoint(point)
	writeAPI.Flush()
	client.Close()
	return nil
}
