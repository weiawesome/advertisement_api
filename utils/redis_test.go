/*
There is a redis client instance.
If program start, it will initialize the instance and try to connect redis database.
Furthermore, get function and close function is to get redis client and close redis connection.
*/

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetRedisClient is to get the redis client
func TestGetRedisClient(t *testing.T) {
	t.Run("Right case get redis database", func(t *testing.T) {
		client := GetRedisClient()
		assert.Equal(t, redisClient, client)
	})
}

// CloseRedis is to close redis database connection
func TestCloseRedis(t *testing.T) {
	t.Run("Right case when db is nil", func(t *testing.T) {
		err := CloseRedis()
		assert.Nil(t, err)
	})
}
