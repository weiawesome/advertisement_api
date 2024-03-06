package sql

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/internal/repository/model"
	"advertisement_api/utils"
	"errors"
	"github.com/stretchr/testify/mock"
	"time"
)

const (
	NormalCase    = "Normal case"
	ErrorCase     = "Error case"
	TimeLimitCase = "Time limit case"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) GetAdvertisement(Age int, Country string, Gender string, Platform string) ([]model.Advertisement, error) {
	if Country == NormalCase {
		return []model.Advertisement{{Title: NormalCase}}, nil
	} else if Country == ErrorCase {
		return []model.Advertisement{}, errors.New("error with " + ErrorCase)
	} else if Country == TimeLimitCase {
		time.Sleep(time.Second*time.Duration(utils.GetDefaultTimeLimitSecond()) + time.Second*5)
		return []model.Advertisement{}, errors.New("error with " + ErrorCase)
	}
	return []model.Advertisement{}, nil
}

func (r *RepositoryMock) AddAdvertisement(data advertisement.AddAdvertisementRequest) error {
	if *data.Title == NormalCase {
		return nil
	} else if *data.Title == ErrorCase {
		return errors.New("error with " + ErrorCase)
	}
	return nil
}
