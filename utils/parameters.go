/*
All the system parameter can get by following function
Furthermore, setting default value for all the parameters in start of file
*/

package utils

// default value setting
var (
	// basic system information
	version = "v1"

	// basic parameter setting (according to document regulation)
	countries = []string{"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "CV", "KH", "CM", "CA", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "SZ", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MK", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW"} // countries get from https://zh.wikipedia.org/wiki/ISO_3166-1
	genders   = []string{"M", "F"}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             // genders type
	platforms = []string{"android", "ios", "web"}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              // platforms type

	// limitation of setting
	minAge             = 1                             // minimum of age
	maxAge             = 100                           // maximum of age
	minLimit           = 1                             //minimum of limit
	maxLimit           = 100                           //maximum of limit
	minOffset          = 0                             //minimum of offset
	maxOffset          = maxDurationRequest - minLimit //maximum of offset
	minCacheMinute     = 15                            //minimum of cache time
	maxCacheMinute     = 60                            //maximum of cache time
	maxDailyRequest    = 3000                          //maximum of daily request
	maxDurationRequest = 1000                          //maximum of duration request

	// default value for parameters
	defaultOffset            = 0                      // default value for offset
	defaultLimit             = 5                      // default value for limit
	defaultAge               = -1                     // default value for age
	defaultGender            = "NoConstraintGender"   // default value for gender
	defaultCountry           = "NoConstraintCountry"  // default value for country
	defaultPlatform          = "NoConstraintPlatform" // default value for platform
	defaultTimeLimitSecond   = 1
	defaultForgetMilliSecond = 100

	// map to get the parameter sets content
	countriesMap map[string]bool // countries map for countries
	gendersMap   map[string]bool // genders map for genders
	platformsMap map[string]bool // platforms map for platforms
)

// InitMaps is to initialize the maps
// Make the country, gender and platform fill the map
func InitMaps() {
	countriesMap = map[string]bool{}
	gendersMap = map[string]bool{}
	platformsMap = map[string]bool{}

	for _, country := range countries {
		countriesMap[country] = true
	}
	for _, gender := range genders {
		gendersMap[gender] = true
	}
	for _, platform := range platforms {
		platformsMap[platform] = true
	}
}

// GetVersion is to get the version parameter
func GetVersion() string {
	return version
}

// GetCountriesMap is to get the countries map
func GetCountriesMap() map[string]bool {
	return countriesMap
}

// GetPlatformsMap is to get the platforms map
func GetPlatformsMap() map[string]bool {
	return platformsMap
}

// GetGendersMap is to get the genders map
func GetGendersMap() map[string]bool {
	return gendersMap
}

// GetMinAge is to get the minimum age
func GetMinAge() int {
	return minAge
}

// GetMaxAge is to get the maximum age
func GetMaxAge() int {
	return maxAge
}

// GetMinLimit is to get the minimum limit
func GetMinLimit() int {
	return minLimit
}

// GetMaxLimit is to get the maximum limit
func GetMaxLimit() int {
	return maxLimit
}

// GetMinOffset is to get the minimum offset
func GetMinOffset() int {
	return minOffset
}

// GetMaxOffset is to get the maximum offset
func GetMaxOffset() int {
	return maxOffset
}

// GetMinCacheMinute is to get the minimum cache time
func GetMinCacheMinute() int {
	return minCacheMinute
}

// GetMaxCacheMinute is to get the maximum cache time
func GetMaxCacheMinute() int {
	return maxCacheMinute
}

// GetMaxDailyRequest is to get the maximum daily request
func GetMaxDailyRequest() int64 {
	return int64(maxDailyRequest)
}

// GetMaxDurationRequest is to get the maximum duration request
func GetMaxDurationRequest() int64 {
	return int64(maxDurationRequest)
}

// GetDefaultOffset is to get the default offset
func GetDefaultOffset() int {
	return defaultOffset
}

// GetDefaultLimit is to get the default limit
func GetDefaultLimit() int {
	return defaultLimit
}

// GetDefaultAge is to get the default age
func GetDefaultAge() int {
	return defaultAge
}

// GetDefaultGender is to get the default gender
func GetDefaultGender() string {
	return defaultGender
}

// GetDefaultCountry is to get the default country
func GetDefaultCountry() string {
	return defaultCountry
}

// GetDefaultPlatform is to get the default platform
func GetDefaultPlatform() string {
	return defaultPlatform
}

// GetDefaultTimeLimitSecond is to get the default platform
func GetDefaultTimeLimitSecond() int {
	return defaultTimeLimitSecond
}

// GetDefaultForgetMilliSecond is to get the default platform
func GetDefaultForgetMilliSecond() int {
	return defaultForgetMilliSecond
}
