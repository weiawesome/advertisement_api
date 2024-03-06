package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewInfluxDBHook(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		hook := NewInfluxDBHook()
		assert.NotNil(t, hook)
	})
}

func TestRun(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		hook := NewInfluxDBHook()
		hook.Run(logger.Info(), logger.GetLevel(), "test log")
	})
	t.Run("Case error", func(t *testing.T) {
		token := EnvInfluxDbToken()
		err := os.Setenv("INFLUXDB_TOKEN", "no exist token")
		if err != nil {
			t.Errorf("error to set environment " + err.Error())
			return
		}
		hook := NewInfluxDBHook()
		hook.Run(logger.Info(), logger.GetLevel(), "test log")
		err = os.Setenv("INFLUXDB_TOKEN", token)
		if err != nil {
			t.Errorf("error to set environment " + err.Error())
			return
		}
	})
}
