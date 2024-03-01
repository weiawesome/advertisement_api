/*
The schema design for the platform condition.
*/

package model

// PlatformCondition is schema for the platform condition
type PlatformCondition struct {
	AdvertisementId int    `gorm:"primaryKey;foreignKey"`             // advertisement_id is primary key also foreign key.
	PlatformCode    string `gorm:"primaryKey;Index;type:VARCHAR(10)"` // platform_code store in type varchar(10) and is primary key.
}
