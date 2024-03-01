/*
The schema design for the country condition.
*/

package model

// CountryCondition is schema for the country condition
type CountryCondition struct {
	AdvertisementId int    `gorm:"primaryKey;foreignKey"`             // advertisement_id is primary key also foreign key.
	CountryISO      string `gorm:"primaryKey;Index;type:VARCHAR(10)"` // country_iso store in type varchar(10) and is primary key.
}
