package routes

import (
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitRoutes(t *testing.T) {
	t.Run("Right case", func(t *testing.T) {
		sqlRepo := new(sql.RepositoryMock)
		redisRepo := new(redis.RepositoryMock)
		r := InitRoutes(sqlRepo, redisRepo)
		assert.NotNil(t, r)
	})
}
