package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRedisClient(t *testing.T) {
	t.Run("Right case get redis database", func(t *testing.T) {
		client := GetRedisClient()
		assert.Equal(t, redisClient, client)
	})
}

func TestCloseRedis(t *testing.T) {
	t.Run("Right case when db is nil", func(t *testing.T) {
		err := CloseRedis()
		assert.Nil(t, err)
	})
}
