package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInfluxDBClient(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		client := NewInfluxDBClient()
		assert.Equal(t, influxDBClient, client)
	})
}
