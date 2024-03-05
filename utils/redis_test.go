package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitRedis(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		err := InitRedis()
		assert.Nil(t, err)
	})
	t.Run("Case error", func(t *testing.T) {
		redisDb := EnvRedisDb()
		err := os.Setenv("REDIS_DB", "Error case")
		if err != nil {
			t.Errorf("error to set redis db environment " + err.Error())
			return
		}
		err = InitRedis()
		assert.NotNil(t, err)
		err = os.Setenv("REDIS_DB", redisDb)
		if err != nil {
			t.Errorf("error to set redis db environment " + err.Error())
			return
		}
	})
}

func TestGetRedisClient(t *testing.T) {
	t.Run("Right case get redis database", func(t *testing.T) {
		client := GetRedisClient()
		assert.Equal(t, redisClient, client)
	})
}

func TestCloseRedis(t *testing.T) {
	t.Run("Right case when db is nil", func(t *testing.T) {
		redisClient = nil
		err := CloseRedis()
		assert.Nil(t, err)
	})
	t.Run("Right case when db is nil", func(t *testing.T) {
		err := InitRedis()
		if err != nil {
			t.Errorf("error to initalize redis client " + err.Error())
			return
		}
		err = CloseRedis()
		assert.Nil(t, err)
	})
}
