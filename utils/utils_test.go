/*
Some tools including getting today, getting location dsn setting.
*/

package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
	"time"
)

// GetToday is a function to get today's date
func TestGetToday(t *testing.T) {
	t.Run("Case get today", func(t *testing.T) {
		location, _ := time.LoadLocation(EnvLocation())
		now := time.Now().In(location)
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
		assert.Equal(t, today, GetToday())
	})
	t.Run("Case get today with not exist location", func(t *testing.T) {
		err := os.Setenv("LOCATION", "Taiwan")
		if err != nil {
			t.Errorf("set environment error " + err.Error())
			return
		}
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		assert.Equal(t, today, GetToday())
		err = os.Unsetenv("LOCATION")
		if err != nil {
			t.Errorf("unset environment error " + err.Error())
			return
		}
	})
}

// GetSqlDsnLocation is to edit SQL DSN setting
func TestGetSqlDsnLocation(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		location := EnvLocation()
		result := strings.Replace(location, "/", "%2F", 1)

		assert.Equal(t, result, GetSqlDsnLocation())
	})
	t.Run("Case right with not exist location", func(t *testing.T) {
		err := os.Setenv("LOCATION", "Taiwan")
		if err != nil {
			t.Errorf("set environment error " + err.Error())
			return
		}

		assert.Equal(t, localLocation, GetSqlDsnLocation())
		err = os.Unsetenv("LOCATION")
		if err != nil {
			t.Errorf("unset environment error " + err.Error())
			return
		}
	})
}
