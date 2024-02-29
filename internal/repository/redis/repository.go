package redis

import (
	"advertisement_api/utils"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	client *redis.Client
}

func NewRepository() *Repository {
	return &Repository{client: utils.GetRedisClient()}
}
