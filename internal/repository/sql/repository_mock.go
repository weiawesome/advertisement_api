/*
The repository for the sql.
*/

package sql

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/internal/repository/model"
	"errors"
	"github.com/stretchr/testify/mock"
)

const (
	NormalCase = "Normal case"
	ErrorCase  = "Error case"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) GetAdvertisement(Age int, Country string, Gender string, Platform string, Offset int, Limit int) ([]model.Advertisement, error) {
	if Country == NormalCase {
		return []model.Advertisement{{Title: NormalCase}}, nil
	} else if Country == ErrorCase {
		return []model.Advertisement{}, errors.New("error with " + ErrorCase)
	}
	return []model.Advertisement{}, errors.New("error with " + ErrorCase)
}

func (r *RepositoryMock) AddAdvertisement(data advertisement.AddAdvertisementRequest) error {
	if *data.Title == NormalCase {
		return nil
	} else if *data.Title == ErrorCase {
		return errors.New("error with " + ErrorCase)
	}
	return errors.New("error with " + ErrorCase)
}
