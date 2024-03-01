/*
All the environment variable can get by following function
Furthermore, setting default value for all the variable in start of file
*/

package utils

import "os"

// default value setting
var (
	defaultSqlHost     = "localhost"       // default sql host
	defaultSqlPort     = "3306"            // default sql port
	defaultSqlDb       = "DefaultDb"       // default sql db name
	defaultSqlUser     = "DefaultUser"     // default sql user
	defaultSqlPassword = "DefaultPassword" // default sql password

	defaultRedisHost     = "localhost"       // default redis host
	defaultRedisPort     = "6379"            // default redis port
	defaultRedisPassword = "DefaultPassword" // default redis password
	defaultRedisDb       = "0"               // default redis db

	defaultInfluxDbHost   = "localhost"     // default influxdb host
	defaultInfluxDbPort   = "8086"          // default influxdb port
	defaultInfluxDbToken  = "DefaultToken"  // default influxdb token
	defaultInfluxDbOrg    = "DefaultOrg"    // default influxdb organization
	defaultInfluxDbBucket = "DefaultBucket" // default influxdb bucket

	defaultLocation = "Asia/Taipei" // default location
)

// EnvMySqlAddress is to get the environment variable of MYSQL_HOST and MYSQL_PORT
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

// EnvMySqlDb is to get the environment variable of MYSQL_DB
func EnvMySqlDb() string {
	var dbName string
	if dbName = os.Getenv("MYSQL_DB"); len(dbName) == 0 {
		dbName = defaultSqlDb
	}
	return dbName
}

// EnvMySqlUser is to get the environment variable of MYSQL_USER
func EnvMySqlUser() string {
	var user string
	if user = os.Getenv("MYSQL_USER"); len(user) == 0 {
		user = defaultSqlUser
	}
	return user
}

// EnvMySqlPassword is to get the environment variable of MYSQL_PASSWORD
func EnvMySqlPassword() string {
	var password string
	if password = os.Getenv("MYSQL_PASSWORD"); len(password) == 0 {
		password = defaultSqlPassword
	}
	return password
}

// EnvRedisAddress is to get the environment variable of REDIS_HOST and REDIS_PORT
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

// EnvRedisPassword is to get the environment variable of REDIS_PASSWORD
func EnvRedisPassword() string {
	var password string
	if password = os.Getenv("REDIS_PASSWORD"); len(password) == 0 {
		password = defaultRedisPassword
	}
	return password
}

// EnvRedisDb is to get the environment variable of REDIS_DB
func EnvRedisDb() string {
	var db string
	if db = os.Getenv("REDIS_DB"); len(db) == 0 {
		db = defaultRedisDb
	}
	return db
}

// EnvInfluxDbAddress is to get the environment variable of INFLUXDB_HOST and INFLUXDB_PORT
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

// EnvInfluxDbToken is to get the environment variable of INFLUXDB_TOKEN
func EnvInfluxDbToken() string {
	var token string
	if token = os.Getenv("INFLUXDB_TOKEN"); len(token) == 0 {
		token = defaultInfluxDbToken
	}
	return token
}

// EnvInfluxDbOrg is to get the environment variable of INFLUXDB_ORG
func EnvInfluxDbOrg() string {
	var org string
	if org = os.Getenv("INFLUXDB_ORG"); len(org) == 0 {
		org = defaultInfluxDbOrg
	}
	return org
}

// EnvInfluxDbBucket is to get the environment variable of INFLUXDB_BUCKET
func EnvInfluxDbBucket() string {
	var bucket string
	if bucket = os.Getenv("INFLUXDB_BUCKET"); len(bucket) == 0 {
		bucket = defaultInfluxDbBucket
	}
	return bucket
}

// EnvLocation is to get the environment variable of LOCATION
func EnvLocation() string {
	var location string
	if location = os.Getenv("LOCATION"); len(location) == 0 {
		location = defaultLocation
	}
	return location
}
