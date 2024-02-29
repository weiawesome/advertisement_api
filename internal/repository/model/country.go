package model

type CountryCondition struct {
	AdvertisementId int    `gorm:"primaryKey;foreignKey"`
	CountryISO      string `gorm:"primaryKey;Index;type:VARCHAR(10)"`
}
