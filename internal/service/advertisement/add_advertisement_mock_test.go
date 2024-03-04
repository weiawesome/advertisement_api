/*
The service for adding advertisement.
It has sql repository to realize interaction with database.
*/

package advertisement

import (
	advertisementRequest "advertisement_api/api/request/advertisement"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Add is to get the content from handler and add advertisement by sql repository
func TestMockAdd(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		service := AddAdvertisementServiceMock{}
		data := advertisementRequest.AddAdvertisementRequest{
			Title: func(s string) *string { return &s }(NormalTitleMock),
		}
		_, err := service.Add(data)
		assert.Nil(t, err)
	})
	t.Run("Case error with day limit", func(t *testing.T) {
		service := AddAdvertisementServiceMock{}
		data := advertisementRequest.AddAdvertisementRequest{
			Title: func(s string) *string { return &s }(DayLimitTitleMock),
		}
		_, err := service.Add(data)
		assert.NotNil(t, err)
		assert.Equal(t, DayLimitTitleMock, err.Error())
	})
	t.Run("Case error with duration limit", func(t *testing.T) {
		service := AddAdvertisementServiceMock{}
		data := advertisementRequest.AddAdvertisementRequest{
			Title: func(s string) *string { return &s }(DurationLimitTitleMock),
		}
		_, err := service.Add(data)
		assert.NotNil(t, err)
		assert.Equal(t, DurationLimitTitleMock, err.Error())
	})
	t.Run("Case error with unknown", func(t *testing.T) {
		service := AddAdvertisementServiceMock{}
		data := advertisementRequest.AddAdvertisementRequest{
			Title: func(s string) *string { return &s }(UnknownErrorTitleMock),
		}
		_, err := service.Add(data)
		assert.NotNil(t, err)
		assert.Equal(t, UnknownErrorTitleMock, err.Error())
	})
	t.Run("Case error with default", func(t *testing.T) {
		service := AddAdvertisementServiceMock{}
		data := advertisementRequest.AddAdvertisementRequest{
			Title: func(s string) *string { return &s }(""),
		}
		_, err := service.Add(data)
		assert.NotNil(t, err)
		assert.Equal(t, UnknownErrorTitleMock, err.Error())
	})
}
