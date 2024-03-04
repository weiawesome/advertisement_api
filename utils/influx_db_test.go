/*
There an influxdb client instance.
Furthermore, there is a constructor for the influxdb connection.
*/

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// NewInfluxDBClient is a constructor for the influxdb client
func TestNewInfluxDBClient(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		client := NewInfluxDBClient()
		assert.Equal(t, influxDBClient, client)
	})
}
