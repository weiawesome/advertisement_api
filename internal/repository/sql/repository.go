/*
The repository for the sql.
*/

package sql

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/internal/repository/model"
	"advertisement_api/utils"
	"gorm.io/gorm"
)

// Repository is the interface of sql repository
type Repository interface {
	GetAdvertisement(Age int, Country string, Gender string, Platform string) ([]model.Advertisement, error)
	AddAdvertisement(data advertisement.AddAdvertisementRequest) error
}

// repository is the structure of the sql repository
type repository struct {
	db      *gorm.DB // database connection by gorm
	dbSlave *gorm.DB // database connection by gorm
}

// NewRepository is the constructor for the sql repository
func NewRepository() Repository {
	// return a new sql repository with the sql database client
	return &repository{db: utils.GetDB(), dbSlave: utils.GetDBSalve()}
}
