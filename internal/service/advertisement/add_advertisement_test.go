/*
The service for adding advertisement.
It has sql repository to realize interaction with database.
*/

package advertisement

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/internal/repository/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAddAdvertisementService(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		service := NewAddAdvertisementService(&sqlRepository)
		assert.NotNil(t, service)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Case right", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		service := addService{SqlRepository: &sqlRepository}

		_, err := service.Add(advertisement.AddAdvertisementRequest{Title: func(s string) *string { return &s }(sql.NormalCase)})
		assert.Equal(t, nil, err)
	})
	t.Run("Case error", func(t *testing.T) {
		sqlRepository := sql.RepositoryMock{}
		service := addService{SqlRepository: &sqlRepository}

		_, err := service.Add(advertisement.AddAdvertisementRequest{Title: func(s string) *string { return &s }(sql.ErrorCase)})
		assert.Equal(t, "error with "+sql.ErrorCase, err.Error())
	})
}
