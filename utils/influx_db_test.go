package utils

import (
	"github.com/stretchr/testify/assert"
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
}
