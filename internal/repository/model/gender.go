/*
The schema design for the gender condition.
*/

package model

// GenderCondition is schema for the gender condition
type GenderCondition struct {
	AdvertisementId int    `gorm:"primaryKey;foreignKey"`             // advertisement_id is primary key also foreign key.
	GenderCode      string `gorm:"primaryKey;Index;type:VARCHAR(10)"` // gender_code store in type varchar(10) and is primary key.
}
