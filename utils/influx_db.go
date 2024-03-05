/*
A hook for zerolog to make it available to log information into influxdb.
*/

package utils

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/rs/zerolog"
	"time"
)

// InfluxDBHook is the hook structure for zerolog
type InfluxDBHook struct {
	Client influxdb2.Client
}

// NewInfluxDBHook is a constructor for the influxdb hook
func NewInfluxDBHook() InfluxDBHook {
	// get the address and token from environment
	url := EnvInfluxDbAddress()
	token := EnvInfluxDbToken()

	// try to connect the influxdb
	hook := InfluxDBHook{influxdb2.NewClient(url, token)}
	return hook
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
	writeAPI := h.Client.WriteAPIBlocking(org, bucket)
	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		LogError(err.Error())
		return
	}
}
