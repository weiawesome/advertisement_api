/*
All the environment variable can get by following function
Furthermore, setting default value for all the variable in start of file
*/

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// EnvMySqlAddress is to get the environment variable of MYSQL_HOST and MYSQL_PORT
func TestEnvMySqlAddress(t *testing.T) {
	t.Run("Case of EnvMySqlAddress", func(t *testing.T) {
		assert.Equal(t, EnvMySqlAddress(), defaultSqlHost+":"+defaultSqlPort)
	})
}

// EnvMySqlDb is to get the environment variable of MYSQL_DB
func TestEnvMySqlDb(t *testing.T) {
	t.Run("Case of EnvMySqlDb", func(t *testing.T) {
		assert.Equal(t, EnvMySqlDb(), defaultSqlDb)
	})
}

// EnvMySqlUser is to get the environment variable of MYSQL_USER
func TestEnvMySqlUser(t *testing.T) {
	t.Run("Case of EnvMySqlUser", func(t *testing.T) {
		assert.Equal(t, EnvMySqlUser(), defaultSqlUser)
	})
}

// EnvMySqlPassword is to get the environment variable of MYSQL_PASSWORD
func TestEnvMySqlPassword(t *testing.T) {
	t.Run("Case of EnvMySqlPassword", func(t *testing.T) {
		assert.Equal(t, EnvMySqlPassword(), defaultSqlPassword)
	})
}

// EnvRedisAddress is to get the environment variable of REDIS_HOST and REDIS_PORT
func TestEnvRedisAddress(t *testing.T) {
	t.Run("Case of EnvRedisAddress", func(t *testing.T) {
		assert.Equal(t, EnvRedisAddress(), defaultRedisHost+":"+defaultRedisPort)
	})
}

// EnvRedisPassword is to get the environment variable of REDIS_PASSWORD
func TestEnvRedisPassword(t *testing.T) {
	t.Run("Case of EnvRedisPassword", func(t *testing.T) {
		assert.Equal(t, EnvRedisPassword(), defaultRedisPassword)
	})
}

// EnvRedisDb is to get the environment variable of REDIS_DB
func TestEnvRedisDb(t *testing.T) {
	t.Run("Case of EnvRedisDb", func(t *testing.T) {
		assert.Equal(t, EnvRedisDb(), defaultRedisDb)
	})
}

// EnvInfluxDbAddress is to get the environment variable of INFLUXDB_HOST and INFLUXDB_PORT
func TestEnvInfluxDbAddress(t *testing.T) {
	t.Run("Case of EnvInfluxDbAddress", func(t *testing.T) {
		assert.Equal(t, EnvInfluxDbAddress(), "http://"+defaultInfluxDbHost+":"+defaultInfluxDbPort)
	})
}

// EnvInfluxDbToken is to get the environment variable of INFLUXDB_TOKEN
func TestEnvInfluxDbToken(t *testing.T) {
	t.Run("Case of EnvInfluxDbToken", func(t *testing.T) {
		assert.Equal(t, EnvInfluxDbToken(), defaultInfluxDbToken)
	})
}

// EnvInfluxDbOrg is to get the environment variable of INFLUXDB_ORG
func TestEnvInfluxDbOrg(t *testing.T) {
	t.Run("Case of EnvInfluxDbOrg", func(t *testing.T) {
		assert.Equal(t, EnvInfluxDbOrg(), defaultInfluxDbOrg)
	})
}

// EnvInfluxDbBucket is to get the environment variable of INFLUXDB_BUCKET
func TestEnvInfluxDbBucket(t *testing.T) {
	t.Run("Case of EnvInfluxDbBucket", func(t *testing.T) {
		assert.Equal(t, EnvInfluxDbBucket(), defaultInfluxDbBucket)
	})
}

// EnvLocation is to get the environment variable of LOCATION
func TestEnvLocation(t *testing.T) {
	t.Run("Case of EnvLocation", func(t *testing.T) {
		assert.Equal(t, EnvLocation(), defaultLocation)
	})
}
