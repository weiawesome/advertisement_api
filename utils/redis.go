/*
There is a redis client instance.
If program start, it will initialize the instance and try to connect redis database.
Furthermore, get function and close function is to get redis client and close redis connection.
*/

package utils

import (
	"github.com/redis/go-redis/v9"
)

// redis client instance
var redisClient *redis.Client

// InitRedis is to initialize connection to redis database
func InitRedis() error {
	// get the password, address and db from environment
	Password := EnvRedisPassword()
	Address := EnvRedisAddress()
	Db := EnvRedisDb()

	// try to connect with redis database
	opt, err := redis.ParseURL("redis://:" + Password + "@" + Address + "/" + Db)
	if err != nil {
		return err
	}

	// if success to connect redis database, it will make a new a redis client instance
	redisClient = redis.NewClient(opt)
	return nil
}

// GetRedisClient is to get the redis client
func GetRedisClient() *redis.Client {
	return redisClient
}

// CloseRedis is to close redis database connection
func CloseRedis() error {
	if redisClient == nil {
		return nil
	}

	err := redisClient.Close()
	if err != nil {
		return err
	}

	return nil
}
