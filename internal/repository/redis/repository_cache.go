/*
The redis repository's method about cache.
Include of loading and saving cache.
*/

package redis

import (
	"advertisement_api/api/response/advertisement"
	"advertisement_api/utils"
	"context"
	"encoding/json"
	"math/rand"
	"time"
)

// SaveCache is to save the cache value
func (r *Repository) SaveCache(key string, value advertisement.GetAdvertisementResponse) error {
	// get the context background
	ctx := context.Background()

	// make the data into the json
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	// get the minimum and maximum time for cache
	minDuration := utils.GetMinCacheMinute()
	maxDuration := utils.GetMaxCacheMinute()

	// get a random expire time to save cache
	randomDuration := rand.Intn(maxDuration-minDuration+1) + minDuration
	expireDuration := time.Minute * time.Duration(randomDuration)

	// save the cache
	_, err = r.client.Set(ctx, key, string(bytes), expireDuration).Result()
	return err
}

// LoadCache is to load the cache with specified key
func (r *Repository) LoadCache(key string) (advertisement.GetAdvertisementResponse, error) {
	// declare a variable for the response
	var response advertisement.GetAdvertisementResponse

	// get the context background
	ctx := context.Background()

	// try to get the cache value. if failed, it will return error.
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return response, err
	}

	// success to get cache and try to reformat to specific type
	err = json.Unmarshal([]byte(result), &response)
	return response, err
}
