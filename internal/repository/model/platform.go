package model

type PlatformCondition struct {
	AdvertisementId int    `gorm:"primaryKey;foreignKey"`
	PlatformCode    string `gorm:"primaryKey;Index;type:VARCHAR(10)"`
}
