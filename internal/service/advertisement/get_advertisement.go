/*
The service for getting advertisements.
It has sql repository to realize interaction with database.
Furthermore, redis repository to get and set the cache.
*/

package advertisement

import (
	"advertisement_api/api/response/advertisement"
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"advertisement_api/utils"
	"context"
	"time"
)

// ServiceGetAdvertisement is the interface of the getting advertisements service
type ServiceGetAdvertisement interface {
	Get(Key string, Age int, Country string, Gender string, Platform string, Offset int, Limit int) (advertisement.GetAdvertisementResponse, error)
}

// getService is the service to get advertisements
type getService struct {
	SqlRepository   sql.Repository   // the repository to interact with sql database
	RedisRepository redis.Repository // the repository to handle cache
}

// NewGetAdvertisementService is the contractor for the getService
func NewGetAdvertisementService(s sql.Repository, r redis.Repository) ServiceGetAdvertisement {
	return &getService{SqlRepository: s, RedisRepository: r}
}

// Get is to get the content from handler and query advertisements by sql and redis repository
func (m *getService) Get(Key string, Age int, Country string, Gender string, Platform string, Offset int, Limit int) (advertisement.GetAdvertisementResponse, error) {
	// get the context and set limit of waiting time
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(utils.GetDefaultTimeLimitSecond())*time.Second)
	defer cancel()

	// using singleflight to handle cache invalid
	ch := utils.GetSingleFlight().DoChan(Key, func() (interface{}, error) {
		// declare the variable for the result
		result := advertisement.GetAdvertisementResponse{}

		cacheKey := utils.GetCacheKey(Age, Country, Gender, Platform)

		// get cache from redis repository
		response, err := m.RedisRepository.LoadCache(cacheKey)

		// set a goroutine to forget the key in specific time
		go func() {
			time.Sleep(time.Duration(utils.GetDefaultForgetMilliSecond()) * time.Millisecond)
			utils.GetSingleFlight().Forget(Key)
		}()

		if err == nil {
			// if the result exist then fill it
			if Offset < len(response.Items) {
				result.Items = response.Items[Offset : Offset+min(len(response.Items)-Offset, Limit)]
			}
			// cache hit then return the cache result immediately
			go utils.LogInfo("Cache hit")
			return result, nil
		} else {
			// cache miss
			go utils.LogInfo("Cache miss")

			// start to find the data using sql repository
			advertisements, err := m.SqlRepository.GetAdvertisement(Age, Country, Gender, Platform)

			if err != nil {
				// error to find the advertisements
				return result, err
			}

			// declare the all result variable
			allResult := advertisement.GetAdvertisementResponse{}

			// enumerate it and fill into the result
			allResult.Items = make([]advertisement.Item, len(advertisements))
			for i, ad := range advertisements {
				allResult.Items[i] = advertisement.Item{Title: ad.Title, EndAt: ad.EndAt.UTC()}
			}

			// if the result exist then fill it
			if Offset < len(allResult.Items) {
				result.Items = allResult.Items[Offset : Offset+min(len(allResult.Items)-Offset, Limit)]
			}

			// use another goroutine to set the cache. if failed, then log its error.
			go func() {
				err := m.RedisRepository.SaveCache(cacheKey, allResult)
				if err != nil {
					utils.LogError(err.Error())
				}
			}()

			// return the result
			return result, nil
		}
	})

	// waiting from the singleflight's channel result
	select {
	case <-ctx.Done():
		// when the context's limit time is exceed then return the error
		return advertisement.GetAdvertisementResponse{}, ctx.Err()
	case ret := <-ch:
		// return the result from the channel
		return ret.Val.(advertisement.GetAdvertisementResponse), ret.Err
	}
}
