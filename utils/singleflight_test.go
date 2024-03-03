/*
The singleflight is to handle with cache-invalid.
There is a singleflight instance. If program start, it will initialize the instance.
There is a get function making other function to use it.
*/

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// InitSingleFlight is the function of initialization singleflight
func TestInitSingleFlight(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		InitSingleFlight()
		assert.NotNil(t, gsf)
	})
}

// GetSingleFlight is to get the singleflight instance
func TestGetSingleFlight(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		assert.Equal(t, gsf, GetSingleFlight())
	})
}
