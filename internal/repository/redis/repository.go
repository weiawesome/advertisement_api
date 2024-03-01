/*
The repository for the redis.
*/

package redis

import (
	"advertisement_api/utils"
	"github.com/redis/go-redis/v9"
)

// Repository is the structure of the redis repository
type Repository struct {
	client *redis.Client // redis client for the redis database connection
}

// NewRepository is the constructor for the redis repository
func NewRepository() *Repository {
	// return a new redis repository with the redis client
	return &Repository{client: utils.GetRedisClient()}
}
