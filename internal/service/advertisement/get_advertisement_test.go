/*
The service for getting advertisements.
It has sql repository to realize interaction with database.
Furthermore, redis repository to get and set the cache.
*/

package advertisement

import (
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"advertisement_api/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Get is to get the content from handler and query advertisements by sql and redis repository
func TestGet(t *testing.T) {
	utils.InitSingleFlight()
	t.Run("Case right (Cache hit)", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		result, err := service.Get(redis.CacheHitCase, 0, sql.NormalCase, "M", "ios", 0, 10)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(result.Items))
		assert.Equal(t, redis.CacheResult, result.Items[0].Title)
	})
	t.Run("Case right (Cache miss)", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		result, err := service.Get(redis.CacheMissCase, 0, sql.NormalCase, "M", "ios", 0, 10)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(result.Items))
		assert.Equal(t, sql.NormalCase, result.Items[0].Title)
	})
	t.Run("Case error", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		_, err := service.Get(redis.CacheMissCase, 0, sql.ErrorCase, "M", "ios", 0, 10)
		assert.Equal(t, "error with "+sql.ErrorCase, err.Error())
	})
	t.Run("Case error", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		_, err := service.Get(redis.CacheNormalCase, 0, sql.ErrorCase, "M", "ios", 0, 10)
		assert.Equal(t, "error with "+sql.ErrorCase, err.Error())
	})
}
