package model

type GenderCondition struct {
	AdvertisementId int    `gorm:"primaryKey;foreignKey"`
	GenderCode      string `gorm:"primaryKey;Index;type:VARCHAR(10)"`
}
