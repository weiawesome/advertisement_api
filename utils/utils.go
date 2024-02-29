package utils

import "time"

func GetToday() time.Time {
	location, err := time.LoadLocation(EnvLocation())
	if err != nil {
		now := time.Now()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
		return today
	}
	now := time.Now().In(location)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	return today
}
