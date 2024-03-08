package advertisement

import (
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"advertisement_api/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGetAdvertisementService(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := NewGetAdvertisementService(&sqlRepository, &redisRepository)
		assert.NotNil(t, service)
	})
}

func TestGet(t *testing.T) {
	utils.InitSingleFlight()
	t.Run("Case right (Cache hit)", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		result, err := service.Get("", 0, sql.NormalCase, redis.CacheHitCase, redis.CacheHitCase, 0, 10)

		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(result.Items))
		assert.Equal(t, redis.CacheResult, result.Items[0].Title)
	})
	t.Run("Case right (Cache hit)", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		result, err := service.Get("", 0, sql.NormalCase, redis.CacheHitCase, redis.CacheHitCase, 10, 10)

		assert.Equal(t, nil, err)
		assert.Equal(t, 0, len(result.Items))
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
	t.Run("Case right (Cache miss)", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		result, err := service.Get(redis.CacheMissCase, 0, sql.NormalCase, "M", "ios", 10, 10)
		assert.Equal(t, nil, err)
		assert.Equal(t, 0, len(result.Items))
	})
	t.Run("Case error", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		_, err := service.Get(redis.CacheMissCase, 0, sql.ErrorCase, redis.CacheMissCase, redis.CacheMissCase, 0, 10)
		assert.Equal(t, "error with "+sql.ErrorCase, err.Error())
	})
	t.Run("Case error with cache write back error", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		result, err := service.Get(redis.CacheNormalCase, 0, sql.NormalCase, redis.CacheMissCase, redis.CacheMissCase, 0, 10)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(result.Items))
		assert.Equal(t, sql.NormalCase, result.Items[0].Title)
	})
	t.Run("Case error with time limit", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		redisRepository := redis.RepositoryMock{}
		service := getService{SqlRepository: &sqlRepository, RedisRepository: &redisRepository}

		_, err := service.Get(redis.CacheNormalCase, 0, sql.TimeLimitCase, "M", "ios", 0, 10)
		assert.NotNil(t, err)
	})
}
