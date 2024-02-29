package utils

import "os"

var (
	defaultSqlHost     = "localhost"
	defaultSqlPort     = "3306"
	defaultSqlDb       = "DefaultDb"
	defaultSqlUser     = "DefaultUser"
	defaultSqlPassword = "DefaultPassword"

	defaultRedisHost     = "localhost"
	defaultRedisPort     = "6379"
	defaultRedisPassword = "DefaultPassword"
	defaultRedisDb       = "DefaultDb"

	defaultInfluxDbHost   = "localhost"
	defaultInfluxDbPort   = "8086"
	defaultInfluxDbToken  = "DefaultToken"
	defaultInfluxDbOrg    = "DefaultOrg"
	defaultInfluxDbBucket = "DefaultBucket"

	defaultLocation = "Asia/Taipei"
)

func EnvMySqlAddress() string {
	var ip string
	var port string
	if ip = os.Getenv("MYSQL_HOST"); len(ip) == 0 {
		ip = defaultSqlHost
	}
	if port = os.Getenv("MYSQL_PORT"); len(port) == 0 {
		port = defaultSqlPort
	}
	return ip + ":" + port
}
func EnvMySqlDb() string {
	var dbName string
	if dbName = os.Getenv("MYSQL_DB"); len(dbName) == 0 {
		dbName = defaultSqlDb
	}
	return dbName
}
func EnvMySqlUser() string {
	var user string
	if user = os.Getenv("MYSQL_USER"); len(user) == 0 {
		user = defaultSqlUser
	}
	return user
}
func EnvMySqlPassword() string {
	var password string
	if password = os.Getenv("MYSQL_PASSWORD"); len(password) == 0 {
		password = defaultSqlPassword
	}
	return password
}

func EnvRedisAddress() string {
	var ip string
	var port string
	if ip = os.Getenv("REDIS_HOST"); len(ip) == 0 {
		ip = defaultRedisHost
	}
	if port = os.Getenv("REDIS_PORT"); len(port) == 0 {
		port = defaultRedisPort
	}
	return ip + ":" + port
}
func EnvRedisPassword() string {
	var password string
	if password = os.Getenv("REDIS_PASSWORD"); len(password) == 0 {
		password = defaultRedisPassword
	}
	return password
}
func EnvRedisDb() string {
	var db string
	if db = os.Getenv("REDIS_DB"); len(db) == 0 {
		db = defaultRedisDb
	}
	return db
}

func EnvInfluxDbAddress() string {
	var host string
	var port string
	if host = os.Getenv("INFLUXDB_HOST"); len(host) == 0 {
		host = "http://" + defaultInfluxDbHost
	} else {
		host = "http://" + host
	}
	if port = os.Getenv("INFLUXDB_PORT"); len(port) == 0 {
		port = defaultInfluxDbPort
	}

	return host + ":" + port
}
func EnvInfluxDbToken() string {
	var token string
	if token = os.Getenv("INFLUXDB_TOKEN"); len(token) == 0 {
		token = defaultInfluxDbToken
	}
	return token
}
func EnvInfluxDbOrg() string {
	var org string
	if org = os.Getenv("INFLUXDB_ORG"); len(org) == 0 {
		org = defaultInfluxDbOrg
	}
	return org
}
func EnvInfluxDbBucket() string {
	var bucket string
	if bucket = os.Getenv("INFLUXDB_BUCKET"); len(bucket) == 0 {
		bucket = defaultInfluxDbBucket
	}
	return bucket
}

func EnvLocation() string {
	var location string
	if location = os.Getenv("LOCATION"); len(location) == 0 {
		location = defaultLocation
	}
	return location
}
