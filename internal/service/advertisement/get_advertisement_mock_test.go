package advertisement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockGet(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		service := GetAdvertisementServiceMock{}
		_, err := service.Get("", 10, NormalCountryMock, "", "", 0, 0)
		assert.Nil(t, err)
	})
	t.Run("Case error", func(t *testing.T) {
		service := GetAdvertisementServiceMock{}
		_, err := service.Get("", 10, UnknownErrorCountryMock, "", "", 0, 0)
		assert.NotNil(t, err)
		assert.Equal(t, UnknownErrorCountryMock, err.Error())
	})
}
