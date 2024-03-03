/*
The repository for the redis.
*/

package redis

import (
	"advertisement_api/api/response/advertisement"
	"errors"
	"github.com/stretchr/testify/mock"
)

const (
	CacheHitCase    = "Cache hit"
	CacheMissCase   = "Cache miss"
	CacheNormalCase = "Cache normal"
	CacheResult     = "Cache result"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) SaveCache(key string, value advertisement.GetAdvertisementResponse) error {
	if key == CacheMissCase {
		return errors.New("error with " + CacheMissCase)
	}
	return nil
}

func (r *RepositoryMock) LoadCache(key string) (advertisement.GetAdvertisementResponse, error) {
	if key == CacheHitCase {
		return advertisement.GetAdvertisementResponse{Items: []advertisement.Item{{Title: CacheResult}}}, nil
	} else if key == CacheMissCase {
		return advertisement.GetAdvertisementResponse{}, errors.New("error with " + CacheMissCase)
	}
	return advertisement.GetAdvertisementResponse{}, errors.New("error with " + CacheMissCase)
}
