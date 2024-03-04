package advertisement

import (
	advertisementRequest "advertisement_api/api/request/advertisement"
	advertisementResponse "advertisement_api/api/response/advertisement"
	"advertisement_api/api/response/failure"
	"github.com/stretchr/testify/mock"
)

const (
	NormalTitleMock        = "Advertisement's title"
	DayLimitTitleMock      = "Exceed day limitation"
	DurationLimitTitleMock = "Exceed duration limitation"
	UnknownErrorTitleMock  = "Unknown error"
)

type AddAdvertisementServiceMock struct {
	mock.Mock
}

func (m *AddAdvertisementServiceMock) Add(data advertisementRequest.AddAdvertisementRequest) (advertisementResponse.AddAdvertisementResponse, error) {
	var result advertisementResponse.AddAdvertisementResponse
	if *data.Title == NormalTitleMock {
		return result, nil
	} else if *data.Title == DayLimitTitleMock {
		return result, failure.DayLimitError{Reason: DayLimitTitleMock}
	} else if *data.Title == DurationLimitTitleMock {
		return result, failure.DurationLimitError{Reason: DurationLimitTitleMock}
	} else if *data.Title == UnknownErrorTitleMock {
		return result, failure.ServerError{Reason: UnknownErrorTitleMock}
	}
	return result, failure.ServerError{Reason: UnknownErrorTitleMock}
}
