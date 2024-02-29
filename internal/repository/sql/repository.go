package sql

import (
	"advertisement_api/utils"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{db: utils.GetDB()}
}
