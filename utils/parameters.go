package utils

var (
	version   = "v1"
	countries = []string{"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "CV", "KH", "CM", "CA", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "SZ", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MK", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW"}
	genders   = []string{"M", "F"}
	platforms = []string{"android", "ios", "web"}

	minAge             = 1
	maxAge             = 100
	minLimit           = 1
	maxLimit           = 100
	minOffset          = 0
	maxDailyRequest    = 3000
	maxDurationRequest = 1000
	minCacheMinute     = 15
	maxCacheMinute     = 60

	defaultOffset   = 0
	defaultLimit    = 5
	defaultAge      = -1
	defaultGender   = "NoConstraintGender"
	defaultCountry  = "NoConstraintCountry"
	defaultPlatform = "NoConstraintPlatform"

	countriesMap map[string]bool
	gendersMap   map[string]bool
	platformsMap map[string]bool
)

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

func GetVersion() string {
	return version
}

func GetCountries() []string {
	return countries
}
func GetCountriesMap() map[string]bool {
	return countriesMap
}
func GetPlatforms() []string {
	return platforms
}
func GetPlatformsMap() map[string]bool {
	return platformsMap
}
func GetGenders() []string {
	return genders
}
func GetGendersMap() map[string]bool {
	return gendersMap
}
func GetMinAge() int {
	return minAge
}
func GetMaxAge() int {
	return maxAge
}
func GetMinLimit() int {
	return minLimit
}
func GetMaxLimit() int {
	return maxLimit
}
func GetMinOffset() int {
	return minOffset
}
func GetMaxDailyRequest() int64 {
	return int64(maxDailyRequest)
}
func GetMaxDurationRequest() int64 {
	return int64(maxDurationRequest)
}
func GetMinCacheMinute() int {
	return minCacheMinute
}
func GetMaxCacheMinute() int {
	return maxCacheMinute
}

func GetDefaultOffset() int {
	return defaultOffset
}
func GetDefaultLimit() int {
	return defaultLimit
}
func GetDefaultAge() int {
	return defaultAge
}
func GetDefaultGender() string {
	return defaultGender
}
func GetDefaultCountry() string {
	return defaultCountry
}
func GetDefaultPlatform() string {
	return defaultPlatform
}
