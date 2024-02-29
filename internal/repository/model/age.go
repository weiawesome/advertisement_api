package model

type AgeCondition struct {
	AdvertisementId int `gorm:"foreignKey;primaryKey;"`
	AgeStart        int
	AgeEnd          int
}
