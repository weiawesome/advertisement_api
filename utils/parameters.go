package utils

func GetCountries() []string {
	return []string{"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH", "BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF", "BI", "CV", "KH", "CM", "CA", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR", "CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "SZ", "ET", "FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP", "GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE", "IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS", "LR", "LY", "LI", "LT", "LU", "MO", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT", "MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE", "NG", "NU", "NF", "MK", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR", "QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS", "SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SE", "CH", "SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE", "GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW"}
}
func GetCountriesMap() map[string]bool {
	return map[string]bool{"AF": true, "AX": true, "AL": true, "DZ": true, "AS": true, "AD": true, "AO": true, "AI": true, "AQ": true, "AG": true, "AR": true, "AM": true, "AW": true, "AU": true, "AT": true, "AZ": true, "BS": true, "BH": true, "BD": true, "BB": true, "BY": true, "BE": true, "BZ": true, "BJ": true, "BM": true, "BT": true, "BO": true, "BQ": true, "BA": true, "BW": true, "BV": true, "BR": true, "IO": true, "BN": true, "BG": true, "BF": true, "BI": true, "CV": true, "KH": true, "CM": true, "CA": true, "KY": true, "CF": true, "TD": true, "CL": true, "CN": true, "CX": true, "CC": true, "CO": true, "KM": true, "CG": true, "CD": true, "CK": true, "CR": true, "CI": true, "HR": true, "CU": true, "CW": true, "CY": true, "CZ": true, "DK": true, "DJ": true, "DM": true, "DO": true, "EC": true, "EG": true, "SV": true, "GQ": true, "ER": true, "EE": true, "SZ": true, "ET": true, "FK": true, "FO": true, "FJ": true, "FI": true, "FR": true, "GF": true, "PF": true, "TF": true, "GA": true, "GM": true, "GE": true, "DE": true, "GH": true, "GI": true, "GR": true, "GL": true, "GD": true, "GP": true, "GU": true, "GT": true, "GG": true, "GN": true, "GW": true, "GY": true, "HT": true, "HM": true, "VA": true, "HN": true, "HK": true, "HU": true, "IS": true, "IN": true, "ID": true, "IR": true, "IQ": true, "IE": true, "IM": true, "IL": true, "IT": true, "JM": true, "JP": true, "JE": true, "JO": true, "KZ": true, "KE": true, "KI": true, "KP": true, "KR": true, "KW": true, "KG": true, "LA": true, "LV": true, "LB": true, "LS": true, "LR": true, "LY": true, "LI": true, "LT": true, "LU": true, "MO": true, "MG": true, "MW": true, "MY": true, "MV": true, "ML": true, "MT": true, "MH": true, "MQ": true, "MR": true, "MU": true, "YT": true, "MX": true, "FM": true, "MD": true, "MC": true, "MN": true, "ME": true, "MS": true, "MA": true, "MZ": true, "MM": true, "NA": true, "NR": true, "NP": true, "NL": true, "NC": true, "NZ": true, "NI": true, "NE": true, "NG": true, "NU": true, "NF": true, "MK": true, "MP": true, "NO": true, "OM": true, "PK": true, "PW": true, "PS": true, "PA": true, "PG": true, "PY": true, "PE": true, "PH": true, "PN": true, "PL": true, "PT": true, "PR": true, "QA": true, "RE": true, "RO": true, "RU": true, "RW": true, "BL": true, "SH": true, "KN": true, "LC": true, "MF": true, "PM": true, "VC": true, "WS": true, "SM": true, "ST": true, "SA": true, "SN": true, "RS": true, "SC": true, "SL": true, "SG": true, "SX": true, "SK": true, "SI": true, "SB": true, "SO": true, "ZA": true, "GS": true, "SS": true, "ES": true, "LK": true, "SD": true, "SR": true, "SJ": true, "SE": true, "CH": true, "SY": true, "TW": true, "TJ": true, "TZ": true, "TH": true, "TL": true, "TG": true, "TK": true, "TO": true, "TT": true, "TN": true, "TR": true, "TM": true, "TC": true, "TV": true, "UG": true, "UA": true, "AE": true, "GB": true, "US": true, "UM": true, "UY": true, "UZ": true, "VU": true, "VE": true, "VN": true, "VG": true, "VI": true, "WF": true, "EH": true, "YE": true, "ZM": true, "ZW": true}
}
func GetPlatforms() []string {
	return []string{"android", "ios", "web"}
}
func GetPlatformsMap() map[string]bool {
	return map[string]bool{"android": true, "ios": true, "web": true}
}
func GetGenders() []string {
	return []string{"M", "F"}
}
func GetGendersMap() map[string]bool {
	return map[string]bool{"M": true, "F": true}
}
func GetMinAge() int {
	return 1
}
func GetMaxAge() int {
	return 100
}
func GetMinLimit() int {
	return 1
}
func GetMaxLimit() int {
	return 100
}

func GetDefaultOffset() int {
	return 0
}
func GetDefaultLimit() int {
	return 5
}
func GetDefaultAge() int {
	return -1
}
func GetDefaultGender() string {
	return "NoConstraintGender"
}
func GetDefaultCountry() string {
	return "NoConstraintCountry"
}
func GetDefaultPlatform() string {
	return "NoConstraintPlatform"
}
