package model

import "time"

type Advertisement struct {
	AdvertisementId int       `gorm:"primaryKey;autoIncrement"`
	Title           string    `gorm:"type:text;not null"`
	StartAt         time.Time `gorm:"index:idx_member,priority:1;not null"`
	EndAt           time.Time `gorm:"index:idx_member,priority:1;not null"`
	CreatedAt       time.Time `gorm:"type:date;index:sort:desc,priority:2;not null"`

	AllAgeCondition      bool `gorm:"type:bool;default:false"`
	AllCountryCondition  bool `gorm:"type:bool;default:false"`
	AllGenderCondition   bool `gorm:"type:bool;default:false"`
	AllPlatformCondition bool `gorm:"type:bool;default:false"`

	Genders   []GenderCondition
	Countries []CountryCondition
	Platforms []PlatformCondition
	Ages      AgeCondition
}
