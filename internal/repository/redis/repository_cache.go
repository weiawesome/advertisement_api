package redis

import (
	"advertisement_api/api/response/advertisement"
	"advertisement_api/utils"
	"context"
	"encoding/json"
	"math/rand"
	"time"
)

func (r *Repository) SaveCache(key string, value advertisement.GetAdvertisementResponse) error {
	ctx := context.Background()
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	minDuration := utils.GetMinCacheMinute()
	maxDuration := utils.GetMaxCacheMinute()

	randomDuration := rand.Intn(maxDuration-minDuration+1) + minDuration
	expireDuration := time.Minute * time.Duration(randomDuration)

	_, err = r.client.Set(ctx, key, string(bytes), expireDuration).Result()
	return err
}
func (r *Repository) LoadCache(key string) (advertisement.GetAdvertisementResponse, error) {
	var response advertisement.GetAdvertisementResponse
	ctx := context.Background()
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return response, err
	}
	err = json.Unmarshal([]byte(result), &response)
	return response, err
}
