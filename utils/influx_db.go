/*
There an influxdb client instance.
Furthermore, there is a constructor for the influxdb connection.
*/

package utils

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog"
	"time"
)

// InfluxDBClient have a influxdb client
type InfluxDBClient struct {
	Client influxdb2.Client
}

// influxdb connection instance
var influxDBClient InfluxDBClient

// NewInfluxDBClient is a constructor for the influxdb client
func NewInfluxDBClient() InfluxDBClient {
	// connection has been existed
	if influxDBClient.Client != nil {
		return influxDBClient
	}

	// get the address and token from environment
	url := EnvInfluxDbAddress()
	token := EnvInfluxDbToken()

	// try to connect the influxdb
	influxDBClient.Client = influxdb2.NewClient(url, token)
	return influxDBClient
}

/*
A hook for zerolog to make it available to log information into influxdb.
*/

// InfluxDBHook is the hook structure for zerolog
type InfluxDBHook struct {
	Client InfluxDBClient
}

// Run is to write the log information
func (h InfluxDBHook) Run(e *zerolog.Event, level zerolog.Level, message string) {
	// set the save point
	point := influxdb2.
		NewPointWithMeasurement("logs").
		SetTime(time.Now()).
		AddTag("level", level.String()).
		AddField("message", message)

	// get the organization and bucket from environment
	org := EnvInfluxDbOrg()
	bucket := EnvInfluxDbBucket()

	// build the write api and write to the influxdb
	writeAPI := h.Client.Client.WriteAPIBlocking(org, bucket)
	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		LogError(err.Error())
		return
	}
}
