/*
There is initialization of the routes.
It will set some setting at first include trusted proxy, template files, groups routes etc.
Then new redis and sql repository. Finally, initialize all routes and return it.
*/

package routes

import (
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

// InitRoutes is to initialize all routes and setting
func TestInitRoutes(t *testing.T) {
	t.Run("Right case", func(t *testing.T) {
		sqlRepo := new(sql.RepositoryMock)
		redisRepo := new(redis.RepositoryMock)
		r := InitRoutes(sqlRepo, redisRepo)
		assert.NotNil(t, r)
	})
}
