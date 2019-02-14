package influx

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"

	queue "github.com/valtyr/shake/pkg/queue"
)

func InfluxPersistor() {
	channel := queue.SubscribeToEvents()

	httpClient, _ := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "shake",
		Password: "yabooty",
	})

	for {
		<-channel

		bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  "shake",
			Precision: "us",
		})

		tags := map[string]string{
			"user_id":   "123",
			"school_id": "1234",
		}
		fields := map[string]interface{}{}

		point, _ := client.NewPoint(
			"cpu",
			tags,
			fields,
			time.Now(),
		)

		bp.AddPoint(point)
		httpClient.Write(bp)
	}
}
