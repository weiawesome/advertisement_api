/*
The repository for the redis.
*/

package redis

import (
	"advertisement_api/api/response/advertisement"
	"advertisement_api/utils"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	SaveCache(key string, value advertisement.GetAdvertisementResponse) error
	LoadCache(key string) (advertisement.GetAdvertisementResponse, error)
}

// Repository is the structure of the redis repository
type repository struct {
	client *redis.Client // redis client for the redis database connection
}

// NewRepository is the constructor for the redis repository
func NewRepository() Repository {
	// return a new redis repository with the redis client
	return &repository{client: utils.GetRedisClient()}
}
