/*
Some tools including getting today, getting location dsn setting.
*/

package utils

import (
	"strings"
	"time"
)

const localLocation = "Local"

// GetToday is a function to get today's date
func GetToday() time.Time {
	// get the location setting from the environment
	location, err := time.LoadLocation(EnvLocation())

	// if failed to load location time, it will use local time to return
	if err != nil {
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		return today
	}

	// return the time with location setting
	now := time.Now().In(location)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	return today
}

// GetSqlDsnLocation is to edit SQL DSN setting
func GetSqlDsnLocation() string {
	// get the location and replace word "/" into "%2F"
	location := EnvLocation()

	if strings.ContainsAny(location, "/") {
		return strings.Replace(location, "/", "%2F", 1)
	}
	return localLocation
}
