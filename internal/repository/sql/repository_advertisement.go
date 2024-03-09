/*
The sql repository's method about advertisement affairs.
Include of getting and saving advertisement.
*/

package sql

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/api/response/failure"
	"advertisement_api/internal/repository/model"
	"advertisement_api/utils"
	"gorm.io/gorm"
	"time"
)

// GetAdvertisement is to get advertisements with different condition like age, gender, platform, country etc.
func (r *repository) GetAdvertisement(Age int, Country string, Gender string, Platform string) ([]model.Advertisement, error) {
	// declare a variable for the result
	var advertisements []model.Advertisement

	// build a translation to make sure correctness in database
	// if the action in transaction fail, then it will roll back the database automatically.
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// find the advertisements that its showing period is in now
		tx = tx.Model(model.Advertisement{}).Where("start_at <= NOW()").Where("end_at >= NOW()")

		// the age equal to default value that mean no-constraint, and it doesn't need to filter more
		if Age != utils.GetDefaultAge() {
			// join age_condition find age between age_start and age_end or all_age_condition is true
			tx = tx.Joins("Ages", r.db.Where("age_start <= ?", Age).Where("age_end >= ?", Age)).Where("age_start <= ?", Age).Where("age_end >= ?", Age).Or("all_age_condition = true")
		}

		// the country equal to default value that mean no-constraint, and it doesn't need to filter more
		if Country != utils.GetDefaultCountry() {
			// join country_condition find country is equal or all_country_condition is true
			tx = tx.Joins("Countries", r.db.Where(&model.CountryCondition{CountryISO: Country})).Where("country_iso = ?", Country).Or("all_country_condition = true")
		}

		// the gender equal to default value that mean no-constraint, and it doesn't need to filter more
		if Gender != utils.GetDefaultGender() {
			// join gender_condition find gender is equal or all_gender_condition is true
			tx = tx.Joins("Genders", r.db.Where(&model.GenderCondition{GenderCode: Gender})).Where("gender_code = ?", Gender).Or("all_gender_condition = true")
		}

		// the platform equal to default value that mean no-constraint, and it doesn't need to filter more
		if Platform != utils.GetDefaultPlatform() {
			// join platform_condition find platform is equal or all_platform_condition is true
			tx = tx.Joins("Platforms", r.db.Where(&model.PlatformCondition{PlatformCode: Platform})).Where("platform_code = ?", Platform).Or("all_platform_condition = true")
		}

		// order by end_at asc and find the result
		tx = tx.Order("end_at ASC").Find(&advertisements)

		// return the query result if success, result's error will be nil
		return tx.Error
	})

	return advertisements, err
}

// AddAdvertisement is to add advertisement with specific information.
func (r *repository) AddAdvertisement(data advertisement.AddAdvertisementRequest) error {
	// start a transaction to ensure database correctness
	// when any of the action fail, the database will roll bak automatically.
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// get today's date
		today := utils.GetToday()

		// check the number of the adding advertisement in today has been exceeded limit or not
		if err := checkDayLimit(tx, today, true); err != nil {
			return err
		}
		// check the number of the adding advertisement in duration has been exceeded limit or not
		if err := checkDurationLimit(tx, *data.StartAt, *data.EndAt, true); err != nil {
			return err
		}

		// make the advertisement with the query's content
		ad := model.Advertisement{Title: *data.Title, StartAt: *data.StartAt, EndAt: *data.EndAt, CreatedAt: today}

		// fill about country condition
		if data.Conditions.Country == nil {
			// country is not constraint
			ad.AllCountryCondition = true
		} else {
			// fill all the query's country
			countries := make([]model.CountryCondition, len(data.Conditions.Country))
			for i, c := range data.Conditions.Country {
				countries[i] = model.CountryCondition{CountryISO: c}
			}
		}

		// fill about gender condition
		if data.Conditions.Gender == nil {
			// gender is not constraint
			ad.AllGenderCondition = true
		} else {
			// fill all the query's gender
			genders := make([]model.GenderCondition, len(data.Conditions.Gender))
			for i, g := range data.Conditions.Gender {
				genders[i] = model.GenderCondition{GenderCode: g}
			}
		}

		// fill about platform condition
		if data.Conditions.Platform == nil {
			// platform is not constraint
			ad.AllPlatformCondition = true
		} else {
			// fill all the query's platform
			platforms := make([]model.PlatformCondition, len(data.Conditions.Platform))
			for i, p := range data.Conditions.Platform {
				platforms[i] = model.PlatformCondition{PlatformCode: p}
			}
		}

		// fill about age condition
		if data.Conditions.AgeStart == nil && data.Conditions.AgeEnd == nil {
			// age is not constraint
			ad.AllAgeCondition = true
		} else {
			// fill the query's start and end of age
			ad.Ages.AgeStart = *data.Conditions.AgeStart
			ad.Ages.AgeEnd = *data.Conditions.AgeEnd
		}

		// create the advertisement. if failed, then return the error.
		if err := tx.Create(&ad).Error; err != nil {
			return err
		}

		// check again about day limit, but the number can equal to the limitation.
		if err := checkDayLimit(tx, today, false); err != nil {
			return err
		}
		// check again about duration limit, but the number can equal to the limitation.
		if err := checkDurationLimit(tx, *data.StartAt, *data.EndAt, false); err != nil {
			return err
		}

		// all is passed then return nil
		return nil
	})

	// return the result
	return err
}

// check today's add advertisement record has been exceeded day limit or not
func checkDayLimit(tx *gorm.DB, today time.Time, checkEqual bool) error {
	// declare a variable for the number of the advertisement adding today
	var dailyRequest int64

	// search the number of the advertisement creating today
	if err := tx.Model(model.Advertisement{}).Where("created_at = ?", today).Count(&dailyRequest).Error; err != nil {
		return err
	}

	// to check the number exceed the limitation or not
	if checkEqual {
		// if check equal case, the number larger than or equal to limitation will return error
		if dailyRequest >= utils.GetMaxDailyRequest() {
			return failure.DayLimitError{Reason: "today add advertisement request has exceeded limit"}
		}
	} else {
		// if not check equal case, the number only larger than limitation will return error
		if dailyRequest > utils.GetMaxDailyRequest() {
			return failure.DayLimitError{Reason: "today add advertisement request has exceeded limit"}
		}
	}

	// return nil mean passing check day limit
	return nil
}

// check in the duration the number of advertisements has been exceeded duration limit or not
func checkDurationLimit(tx *gorm.DB, startAt time.Time, endAt time.Time, checkEqual bool) error {
	// declare a variable for the number of the advertisement in specified time period
	var periodRequest int64

	// search the number of the advertisement in the period
	if err := tx.Model(model.Advertisement{}).Where("start_at <= ?", endAt).Where("end_at >= ?", startAt).Count(&periodRequest).Error; err != nil {
		return err
	}

	// to check the number exceed the limitation or not
	if checkEqual {
		// if check equal case, the number larger than or equal to limitation will return error
		if periodRequest >= utils.GetMaxDurationRequest() {
			return failure.DurationLimitError{Reason: "duration in the request has exceeded limit"}
		}
	} else {
		// if not check equal case, the number only larger than limitation will return error
		if periodRequest > utils.GetMaxDurationRequest() {
			return failure.DurationLimitError{Reason: "duration in the request has exceeded limit"}
		}
	}

	// return nil mean passing check duration limit
	return nil
}
