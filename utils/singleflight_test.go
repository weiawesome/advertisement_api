package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitSingleFlight(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		InitSingleFlight()
		assert.NotNil(t, gsf)
	})
}

func TestGetSingleFlight(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		assert.Equal(t, gsf, GetSingleFlight())
	})
}
