/*
The schema design for the age condition.
*/

package model

// AgeCondition is schema for the age condition
type AgeCondition struct {
	AdvertisementId int `gorm:"foreignKey;primaryKey;"` // advertisement_id is primary key also foreign key and primary key to ensure one-to-one relationship.
	AgeStart        int // age_start is the constraint for the age start
	AgeEnd          int // age_end is the constraint for the age end
}
