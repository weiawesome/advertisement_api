/*
The service for adding advertisement.
It has sql repository to realize interaction with database.
*/

package advertisement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Add is to get the content from handler and add advertisement by sql repository
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
