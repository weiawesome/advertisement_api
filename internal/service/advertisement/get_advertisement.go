package advertisement

import (
	"advertisement_api/api/response/advertisement"
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"advertisement_api/utils"
	"context"
	"time"
)

type ServiceGetAdvertisement struct {
	SqlRepository   sql.Repository
	RedisRepository redis.Repository
}

func (m *ServiceGetAdvertisement) Get(Key string, Age int, Country string, Gender string, Platform string, Offset int, Limit int) (advertisement.GetAdvertisementResponse, error) {
	result := advertisement.GetAdvertisementResponse{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := utils.GetSingleFlight().DoChan(Key, func() (interface{}, error) {
		response, err := m.RedisRepository.LoadCache(Key)

		go func() {
			time.Sleep(100 * time.Millisecond)
			utils.GetSingleFlight().Forget(Key)
		}()
		if err == nil {
			go utils.LogInfo("Hit cache")
			return response, nil
		} else {
			go utils.LogInfo("Not hit cache")
			advertisements, err := m.SqlRepository.GetAdvertisement(Age, Country, Gender, Platform, Offset, Limit)
			if err != nil {
				return nil, err
			}
			for _, ad := range advertisements {
				result.Items = append(result.Items, advertisement.Item{Title: ad.Title, EndAt: ad.EndAt})
			}

			go func() {
				err := m.RedisRepository.SaveCache(Key, result)
				if err != nil {
					utils.LogError(err.Error())
				}
			}()

			return result, nil
		}
	})

	select {
	case <-ctx.Done():
		return result, ctx.Err()
	case ret := <-ch:
		return ret.Val.(advertisement.GetAdvertisementResponse), ret.Err
	}
}
