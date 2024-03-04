package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestGetVersion(t *testing.T) {
	t.Run("Case with Version", func(t *testing.T) {
		assert.Equal(t, version, GetVersion())
	})
}

func TestGetCountriesMap(t *testing.T) {
	t.Run("Case with CountriesMap", func(t *testing.T) {
		assert.Equal(t, countriesMap, GetCountriesMap())
	})
	t.Run("Case with CountriesMap after initialization", func(t *testing.T) {
		InitMaps()
		assert.Equal(t, countriesMap, GetCountriesMap())
	})
}

func TestGetPlatformsMap(t *testing.T) {
	t.Run("Case with PlatformsMap", func(t *testing.T) {
		assert.Equal(t, platformsMap, GetPlatformsMap())
	})
	t.Run("Case with PlatformsMap after initialization", func(t *testing.T) {
		InitMaps()
		assert.Equal(t, platformsMap, GetPlatformsMap())
	})
}

func TestGetGendersMap(t *testing.T) {
	t.Run("Case with GendersMap", func(t *testing.T) {
		assert.Equal(t, gendersMap, GetGendersMap())
	})
	t.Run("Case with GendersMap after initialization", func(t *testing.T) {
		InitMaps()
		assert.Equal(t, gendersMap, GetGendersMap())
	})
}

func TestGetMinAge(t *testing.T) {
	t.Run("Case with MinAge", func(t *testing.T) {
		assert.Equal(t, minAge, GetMinAge())
	})
}

func TestGetMaxAge(t *testing.T) {
	t.Run("Case with MaxAge", func(t *testing.T) {
		assert.Equal(t, maxAge, GetMaxAge())
	})
}

func TestGetMinLimit(t *testing.T) {
	t.Run("Case with MinLimit", func(t *testing.T) {
		assert.Equal(t, minLimit, GetMinLimit())
	})
}

func TestGetMaxLimit(t *testing.T) {
	t.Run("Case with MaxLimit", func(t *testing.T) {
		assert.Equal(t, maxLimit, GetMaxLimit())
	})
}

func TestGetMinOffset(t *testing.T) {
	t.Run("Case with MinOffset", func(t *testing.T) {
		assert.Equal(t, minOffset, GetMinOffset())
	})
}

func TestGetMinCacheMinute(t *testing.T) {
	t.Run("Case with MinCacheMinute", func(t *testing.T) {
		assert.Equal(t, minCacheMinute, GetMinCacheMinute())
	})
}

func TestGetMaxCacheMinute(t *testing.T) {
	t.Run("Case with MaxCacheMinute", func(t *testing.T) {
		assert.Equal(t, maxCacheMinute, GetMaxCacheMinute())
	})
}

func TestGetMaxDailyRequest(t *testing.T) {
	t.Run("Case with MaxDailyRequest", func(t *testing.T) {
		assert.Equal(t, int64(maxDailyRequest), GetMaxDailyRequest())
	})
}

func TestGetMaxDurationRequest(t *testing.T) {
	t.Run("Case with MaxDurationRequest", func(t *testing.T) {
		assert.Equal(t, int64(maxDurationRequest), GetMaxDurationRequest())
	})
}

func TestGetDefaultOffset(t *testing.T) {
	t.Run("Case with DefaultOffset", func(t *testing.T) {
		assert.Equal(t, defaultOffset, GetDefaultOffset())
	})
}

func TestGetDefaultLimit(t *testing.T) {
	t.Run("Case with DefaultLimit", func(t *testing.T) {
		assert.Equal(t, defaultLimit, GetDefaultLimit())
	})
}

func TestGetDefaultAge(t *testing.T) {
	t.Run("Case with DefaultAge", func(t *testing.T) {
		assert.Equal(t, defaultAge, GetDefaultAge())
	})
}

func TestGetDefaultGender(t *testing.T) {
	t.Run("Case with DefaultGender", func(t *testing.T) {
		assert.Equal(t, defaultGender, GetDefaultGender())
	})
}

func TestGetDefaultCountry(t *testing.T) {
	t.Run("Case with DefaultCountry", func(t *testing.T) {
		assert.Equal(t, defaultCountry, GetDefaultCountry())
	})
}

func TestGetDefaultPlatform(t *testing.T) {
	t.Run("Case with DefaultPlatform", func(t *testing.T) {
		assert.Equal(t, defaultPlatform, GetDefaultPlatform())
	})
}

func TestGetDefaultTimeLimitSecond(t *testing.T) {
	t.Run("Case with DefaultTimeLimitSecond", func(t *testing.T) {
		assert.Equal(t, defaultTimeLimitSecond, GetDefaultTimeLimitSecond())
	})
}

func TestGetDefaultForgetMilliSecond(t *testing.T) {
	t.Run("Case with DefaultForgetMilliSecond", func(t *testing.T) {
		assert.Equal(t, defaultForgetMilliSecond, GetDefaultForgetMilliSecond())
	})
}
