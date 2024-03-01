/*
The repository for the sql.
*/

package sql

import (
	"advertisement_api/utils"
	"gorm.io/gorm"
)

// Repository is the structure of the sql repository
type Repository struct {
	db *gorm.DB // database connection by gorm
}

// NewRepository is the constructor for the sql repository
func NewRepository() *Repository {
	// return a new sql repository with the sql database client
	return &Repository{db: utils.GetDB()}
}
