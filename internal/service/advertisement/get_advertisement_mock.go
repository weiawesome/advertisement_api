package advertisement

import (
	advertisementResponse "advertisement_api/api/response/advertisement"
	"advertisement_api/api/response/failure"
	"github.com/stretchr/testify/mock"
)

const (
	NormalCountryMock       = "Normal key"
	UnknownErrorCountryMock = "Unknown error"
)

type GetAdvertisementServiceMock struct {
	mock.Mock
}

func (m *GetAdvertisementServiceMock) Get(Key string, Age int, Country string, Gender string, Platform string, Offset int, Limit int) (advertisementResponse.GetAdvertisementResponse, error) {
	var result advertisementResponse.GetAdvertisementResponse

	if Country == NormalCountryMock {
		return result, nil
	}
	return result, failure.ServerError{Reason: UnknownErrorCountryMock}
}
