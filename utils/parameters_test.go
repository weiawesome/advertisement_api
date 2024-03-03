/*
All the system parameter can get by following function
Furthermore, setting default value for all the parameters in start of file
*/

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// InitMaps is to initialize the maps
// Make the country, gender and platform fill the map
func TestInitMaps(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		InitMaps()
		for _, country := range countries {
			assert.True(t, countriesMap[country])
		}
		for _, gender := range genders {
			assert.True(t, gendersMap[gender])
		}
		for _, platform := range platforms {
			assert.True(t, platformsMap[platform])
		}
	})
}

// GetVersion is to get the version parameter
func TestGetVersion(t *testing.T) {
	t.Run("Case with Version", func(t *testing.T) {
		assert.Equal(t, version, GetVersion())
	})
}

// GetCountriesMap is to get the countries map
func TestGetCountriesMap(t *testing.T) {
	t.Run("Case with CountriesMap", func(t *testing.T) {
		assert.Equal(t, countriesMap, GetCountriesMap())
	})
	t.Run("Case with CountriesMap after initialization", func(t *testing.T) {
		InitMaps()
		assert.Equal(t, countriesMap, GetCountriesMap())
	})
}

// GetPlatformsMap is to get the platforms map
func TestGetPlatformsMap(t *testing.T) {
	t.Run("Case with PlatformsMap", func(t *testing.T) {
		assert.Equal(t, platformsMap, GetPlatformsMap())
	})
	t.Run("Case with PlatformsMap after initialization", func(t *testing.T) {
		InitMaps()
		assert.Equal(t, platformsMap, GetPlatformsMap())
	})
}

// GetGendersMap is to get the genders map
func TestGetGendersMap(t *testing.T) {
	t.Run("Case with GendersMap", func(t *testing.T) {
		assert.Equal(t, gendersMap, GetGendersMap())
	})
	t.Run("Case with GendersMap after initialization", func(t *testing.T) {
		InitMaps()
		assert.Equal(t, gendersMap, GetGendersMap())
	})
}

// GetMinAge is to get the minimum age
func TestGetMinAge(t *testing.T) {
	t.Run("Case with MinAge", func(t *testing.T) {
		assert.Equal(t, minAge, GetMinAge())
	})
}

// GetMaxAge is to get the maximum age
func TestGetMaxAge(t *testing.T) {
	t.Run("Case with MaxAge", func(t *testing.T) {
		assert.Equal(t, maxAge, GetMaxAge())
	})
}

// GetMinLimit is to get the minimum limit
func TestGetMinLimit(t *testing.T) {
	t.Run("Case with MinLimit", func(t *testing.T) {
		assert.Equal(t, minLimit, GetMinLimit())
	})
}

// GetMaxLimit is to get the maximum limit
func TestGetMaxLimit(t *testing.T) {
	t.Run("Case with MaxLimit", func(t *testing.T) {
		assert.Equal(t, maxLimit, GetMaxLimit())
	})
}

// GetMinOffset is to get the minimum offset
func TestGetMinOffset(t *testing.T) {
	t.Run("Case with MinOffset", func(t *testing.T) {
		assert.Equal(t, minOffset, GetMinOffset())
	})
}

// GetMinCacheMinute is to get the minimum cache time
func TestGetMinCacheMinute(t *testing.T) {
	t.Run("Case with MinCacheMinute", func(t *testing.T) {
		assert.Equal(t, minCacheMinute, GetMinCacheMinute())
	})
}

// GetMaxCacheMinute is to get the maximum cache time
func TestGetMaxCacheMinute(t *testing.T) {
	t.Run("Case with MaxCacheMinute", func(t *testing.T) {
		assert.Equal(t, maxCacheMinute, GetMaxCacheMinute())
	})
}

// GetMaxDailyRequest is to get the maximum daily request
func TestGetMaxDailyRequest(t *testing.T) {
	t.Run("Case with MaxDailyRequest", func(t *testing.T) {
		assert.Equal(t, int64(maxDailyRequest), GetMaxDailyRequest())
	})
}

// GetMaxDurationRequest is to get the maximum duration request
func TestGetMaxDurationRequest(t *testing.T) {
	t.Run("Case with MaxDurationRequest", func(t *testing.T) {
		assert.Equal(t, int64(maxDurationRequest), GetMaxDurationRequest())
	})
}

// GetDefaultOffset is to get the default offset
func TestGetDefaultOffset(t *testing.T) {
	t.Run("Case with DefaultOffset", func(t *testing.T) {
		assert.Equal(t, defaultOffset, GetDefaultOffset())
	})
}

// GetDefaultLimit is to get the default limit
func TestGetDefaultLimit(t *testing.T) {
	t.Run("Case with DefaultLimit", func(t *testing.T) {
		assert.Equal(t, defaultLimit, GetDefaultLimit())
	})
}

// GetDefaultAge is to get the default age
func TestGetDefaultAge(t *testing.T) {
	t.Run("Case with DefaultAge", func(t *testing.T) {
		assert.Equal(t, defaultAge, GetDefaultAge())
	})
}

// GetDefaultGender is to get the default gender
func TestGetDefaultGender(t *testing.T) {
	t.Run("Case with DefaultGender", func(t *testing.T) {
		assert.Equal(t, defaultGender, GetDefaultGender())
	})
}

// GetDefaultCountry is to get the default country
func TestGetDefaultCountry(t *testing.T) {
	t.Run("Case with DefaultCountry", func(t *testing.T) {
		assert.Equal(t, defaultCountry, GetDefaultCountry())
	})
}

// GetDefaultPlatform is to get the default platform
func TestGetDefaultPlatform(t *testing.T) {
	t.Run("Case with DefaultPlatform", func(t *testing.T) {
		assert.Equal(t, defaultPlatform, GetDefaultPlatform())
	})
}
