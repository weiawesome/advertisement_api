/*
The schema design for the advertisement.
*/

package model

import "time"

// Advertisement is database schema for the table advertisements
//
//	AdvertisementId is the advertisement's id
//	Title is the advertisement's title
//	StartAt and EndAt is the time of advertisement's start and end
//	CreatedAt is the day of the advertisement created
//	AllCondition mean is not constraint for the condition like age, country, gender, platform.
type Advertisement struct {
	AdvertisementId int       `gorm:"primaryKey;autoIncrement"`                      // advertisement_id is primary key and auto increase.
	Title           string    `gorm:"type:text;not null"`                            // title can't be null and store type with text.
	StartAt         time.Time `gorm:"index:idx_member,priority:1;not null"`          // start_at is index with end_at priority is 1 and can't be null.
	EndAt           time.Time `gorm:"index:idx_member,priority:1;not null"`          // end_at is index with start_at priority is 1 and can't be null.
	CreatedAt       time.Time `gorm:"type:date;index:sort:desc,priority:2;not null"` // created_at is index with priority 2 and store type is date and can't be null.

	AllAgeCondition      bool `gorm:"type:bool;default:false"` // all_age_condition store with type bool and default is false
	AllCountryCondition  bool `gorm:"type:bool;default:false"` // all_country_condition store with type bool and default is false
	AllGenderCondition   bool `gorm:"type:bool;default:false"` // all_gender_condition store with type bool and default is false
	AllPlatformCondition bool `gorm:"type:bool;default:false"` // all_platform_condition store with type bool and default is false

	Genders   []GenderCondition   // one-to-many relationship
	Countries []CountryCondition  // one-to-many relationship
	Platforms []PlatformCondition // one-to-many relationship
	Ages      AgeCondition        // one-to-one relationship
}
