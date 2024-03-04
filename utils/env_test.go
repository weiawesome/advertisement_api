package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvMySqlAddress(t *testing.T) {
	t.Run("Case of EnvMySqlAddress", func(t *testing.T) {
		assert.Equal(t, EnvMySqlAddress(), defaultSqlHost+":"+defaultSqlPort)
	})
}

func TestEnvMySqlDb(t *testing.T) {
	t.Run("Case of EnvMySqlDb", func(t *testing.T) {
		assert.Equal(t, EnvMySqlDb(), defaultSqlDb)
	})
}

func TestEnvMySqlUser(t *testing.T) {
	t.Run("Case of EnvMySqlUser", func(t *testing.T) {
		assert.Equal(t, EnvMySqlUser(), defaultSqlUser)
	})
}

func TestEnvMySqlPassword(t *testing.T) {
	t.Run("Case of EnvMySqlPassword", func(t *testing.T) {
		assert.Equal(t, EnvMySqlPassword(), defaultSqlPassword)
	})
}

func TestEnvRedisAddress(t *testing.T) {
	t.Run("Case of EnvRedisAddress", func(t *testing.T) {
		assert.Equal(t, EnvRedisAddress(), defaultRedisHost+":"+defaultRedisPort)
	})
}

func TestEnvRedisPassword(t *testing.T) {
	t.Run("Case of EnvRedisPassword", func(t *testing.T) {
		assert.Equal(t, EnvRedisPassword(), defaultRedisPassword)
	})
}

func TestEnvRedisDb(t *testing.T) {
	t.Run("Case of EnvRedisDb", func(t *testing.T) {
		assert.Equal(t, EnvRedisDb(), defaultRedisDb)
	})
}

func TestEnvInfluxDbAddress(t *testing.T) {
	t.Run("Case of EnvInfluxDbAddress", func(t *testing.T) {
		assert.Equal(t, EnvInfluxDbAddress(), "http://"+defaultInfluxDbHost+":"+defaultInfluxDbPort)
	})
	t.Run("Case of EnvInfluxDbAddress", func(t *testing.T) {
		hostCase := "Test Influxdb Host"
		err := os.Setenv("INFLUXDB_HOST", hostCase)
		if err != nil {
			t.Errorf("set env for influxdb host error " + err.Error())
			return
		}
		assert.Equal(t, EnvInfluxDbAddress(), "http://"+hostCase+":"+defaultInfluxDbPort)
		err = os.Unsetenv("INFLUXDB_HOST")
		if err != nil {
			t.Errorf("set env for influxdb host error " + err.Error())
			return
		}

	})
}

func TestEnvInfluxDbToken(t *testing.T) {
	t.Run("Case of EnvInfluxDbToken", func(t *testing.T) {
		assert.Equal(t, EnvInfluxDbToken(), defaultInfluxDbToken)
	})
}

func TestEnvInfluxDbOrg(t *testing.T) {
	t.Run("Case of EnvInfluxDbOrg", func(t *testing.T) {
		assert.Equal(t, EnvInfluxDbOrg(), defaultInfluxDbOrg)
	})
}

func TestEnvInfluxDbBucket(t *testing.T) {
	t.Run("Case of EnvInfluxDbBucket", func(t *testing.T) {
		assert.Equal(t, EnvInfluxDbBucket(), defaultInfluxDbBucket)
	})
}

func TestEnvLocation(t *testing.T) {
	t.Run("Case of EnvLocation", func(t *testing.T) {
		assert.Equal(t, EnvLocation(), defaultLocation)
	})
}
