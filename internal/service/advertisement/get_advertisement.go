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

// ServiceGetAdvertisement is the service to get advertisements
type ServiceGetAdvertisement struct {
	SqlRepository   sql.Repository   // the repository to interact with sql database
	RedisRepository redis.Repository // the repository to handle cache
}

// Get is to get the content from handler and query advertisements by sql and redis repository
func (m *ServiceGetAdvertisement) Get(Key string, Age int, Country string, Gender string, Platform string, Offset int, Limit int) (advertisement.GetAdvertisementResponse, error) {
	// get the context and set limit of waiting time
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// using singleflight to handle cache invalid
	ch := utils.GetSingleFlight().DoChan(Key, func() (interface{}, error) {
		// declare the variable for the result
		result := advertisement.GetAdvertisementResponse{}

		// get cache from redis repository
		response, err := m.RedisRepository.LoadCache(Key)

		// set a goroutine to forget the key in 0.1 second
		go func() {
			time.Sleep(100 * time.Millisecond)
			utils.GetSingleFlight().Forget(Key)
		}()

		if err == nil {
			// cache hit then return the cache result immediately
			go utils.LogInfo("Cache hit")
			return response, nil
		} else {
			// cache miss
			go utils.LogInfo("Cache miss")

			// start to find the data using sql repository
			advertisements, err := m.SqlRepository.GetAdvertisement(Age, Country, Gender, Platform, Offset, Limit)

			if err != nil {
				// error to find the advertisements
				return response, err
			}

			// enumerate the result
			for _, ad := range advertisements {
				// add each advertisement's information into result
				result.Items = append(result.Items, advertisement.Item{Title: ad.Title, EndAt: ad.EndAt})
			}

			// use another goroutine to set the cache. if failed, then log its error.
			go func() {
				err := m.RedisRepository.SaveCache(Key, result)
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
