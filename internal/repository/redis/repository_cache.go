package redis

import (
	"advertisement_api/api/response/advertisement"
	"context"
	"encoding/json"
	"time"
)

func (r *Repository) SaveCache(key string, value advertisement.GetAdvertisementResponse) error {
	ctx := context.Background()
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	duration := 5
	expireDuration := time.Minute * time.Duration(duration)
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
