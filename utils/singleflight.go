/*
The singleflight is to handle with cache-invalid.
There is a singleflight instance. If program start, it will initialize the instance.
There is a get function making other function to use it.
*/

package utils

import singleflight "golang.org/x/sync/singleflight"

// the singleflight instance
var gsf *singleflight.Group

// InitSingleFLight is the function of initialization singleflight
func InitSingleFLight() {
	gsf = &singleflight.Group{}
}

// GetSingleFlight is to get the singleflight instance
func GetSingleFlight() *singleflight.Group {
	return gsf
}
