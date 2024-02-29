package utils

import singleflight "golang.org/x/sync/singleflight"

var gsf *singleflight.Group

func InitSingleFLight() {
	gsf = &singleflight.Group{}
}

func GetSingleFlight() *singleflight.Group {
	return gsf
}
