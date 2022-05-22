package pkg

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func NewDbConnection(env Env) (influxdb2.Client, api.WriteAPI) {
	client := influxdb2.NewClientWithOptions(
		env.DbHost,
		env.DbToken,
		influxdb2.DefaultOptions().SetBatchSize(20))
	writeAPI := client.WriteAPI(env.DbOrganization, env.DbBucket)

	return client, writeAPI
}

func SaveTelemetryLapData(api api.WriteAPI, packet TelemetryLapData) {
	p := influxdb2.NewPoint(
		"system",
		map[string]string{
			"game": "f1-2021",
			"type": "lap-data",
		},
		map[string]interface{}{
			"lastTime": packet.LastTime,
		},
		time.Now())

	api.WritePoint(p)
}

func SaveTelemetryCarData(api api.WriteAPI, packet TelemetryCarData) {
	p := influxdb2.NewPoint(
		"system",
		map[string]string{
			"game": "f1-2021",
			"type": "car-telemetry",
		},
		map[string]interface{}{
			"speed":    packet.Speed,
			"throttle": packet.Throttle,
			"brake":    packet.Brake,
			"steering": packet.Steering,
		},
		time.Now())

	api.WritePoint(p)
}
