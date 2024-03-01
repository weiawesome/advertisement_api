package sql

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/api/response/failure"
	"advertisement_api/internal/repository/model"
	"advertisement_api/utils"
	"gorm.io/gorm"
	"time"
)

func (r *Repository) GetAdvertisement(Age int, Country string, Gender string, Platform string, Offset int, Limit int) ([]model.Advertisement, error) {
	var advertisements []model.Advertisement
	err := r.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Model(model.Advertisement{}).Where("start_at <= NOW()").Where("end_at >= NOW()")

		if Age != utils.GetDefaultAge() {
			tx = tx.Joins("Ages", r.db.Where("age_start <= ?", Age).Where("age_end >= ?", Age)).Where("age_start <= ?", Age).Where("age_end >= ?", Age).Or("all_age_condition = true")
		}

		if Country != utils.GetDefaultCountry() {
			tx = tx.Joins("Countries", r.db.Where(&model.CountryCondition{CountryISO: Country})).Where("country_iso = ?", Country).Or("all_country_condition = true")
		}

		if Gender != utils.GetDefaultGender() {
			tx = tx.Joins("Genders", r.db.Where(&model.GenderCondition{GenderCode: Gender})).Where("gender_code = ?", Gender).Or("all_gender_condition = true")
		}

		if Platform != utils.GetDefaultPlatform() {
			tx = tx.Joins("Platforms", r.db.Where(&model.PlatformCondition{PlatformCode: Platform})).Where("platform_code = ?", Platform).Or("all_platform_condition = true")
		}

		tx = tx.Order("end_at DESC").Offset(Offset).Limit(Limit).Find(&advertisements)

		return tx.Error
	})

	return advertisements, err
}

func (r *Repository) AddAdvertisement(data advertisement.AddAdvertisementRequest) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		today := utils.GetToday()

		if err := checkDayLimit(tx, today, true); err != nil {
			return err
		}
		if err := checkDurationLimit(tx, data.StartAt, data.EndAt, true); err != nil {
			return err
		}

		ad := model.Advertisement{Title: *data.Title, StartAt: data.StartAt, EndAt: data.EndAt, CreatedAt: today}

		if data.Conditions.Country == nil {
			ad.AllCountryCondition = true
		} else {
			countries := make([]model.CountryCondition, len(data.Conditions.Country))
			for i, c := range data.Conditions.Country {
				countries[i] = model.CountryCondition{CountryISO: c}
			}
		}

		if data.Conditions.Gender == nil {
			ad.AllGenderCondition = true
		} else {
			genders := make([]model.GenderCondition, len(data.Conditions.Gender))
			for i, g := range data.Conditions.Gender {
				genders[i] = model.GenderCondition{GenderCode: g}
			}
		}

		if data.Conditions.Platform == nil {
			ad.AllPlatformCondition = true
		} else {
			platforms := make([]model.PlatformCondition, len(data.Conditions.Platform))
			for i, p := range data.Conditions.Platform {
				platforms[i] = model.PlatformCondition{PlatformCode: p}
			}
		}

		if data.Conditions.AgeStart == nil && data.Conditions.AgeEnd == nil {
			ad.AllAgeCondition = true
		} else {
			ad.Ages.AgeStart = *data.Conditions.AgeStart
			ad.Ages.AgeEnd = *data.Conditions.AgeEnd
		}

		if err := tx.Create(&ad).Error; err != nil {
			return err
		}
		if err := checkDayLimit(tx, today, false); err != nil {
			return err
		}
		if err := checkDurationLimit(tx, data.StartAt, data.EndAt, false); err != nil {
			return err
		}
		return nil
	})
	return err
}

func checkDayLimit(tx *gorm.DB, today time.Time, checkEqual bool) error {
	var dailyRequest int64
	if err := tx.Model(model.Advertisement{}).Where("created_at = ?", today).Count(&dailyRequest).Error; err != nil {
		return err
	}

	if checkEqual {
		if dailyRequest >= utils.GetMaxDailyRequest() {
			return failure.DayLimitError{Reason: "today add advertisement request has exceeded limit"}
		}
	} else {
		if dailyRequest > utils.GetMaxDailyRequest() {
			return failure.DayLimitError{Reason: "today add advertisement request has exceeded limit"}
		}
	}

	return nil
}
func checkDurationLimit(tx *gorm.DB, startAt time.Time, endAt time.Time, checkEqual bool) error {
	var periodRequest int64
	if err := tx.Model(model.Advertisement{}).Where("start_at <= ?", startAt).Where("end_at >= ?", endAt).Count(&periodRequest).Error; err != nil {
		return err
	}

	if checkEqual {
		if periodRequest >= utils.GetMaxDurationRequest() {
			return failure.DurationLimitError{Reason: "duration in the request has exceeded limit"}
		}
	} else {
		if periodRequest > utils.GetMaxDurationRequest() {
			return failure.DurationLimitError{Reason: "duration in the request has exceeded limit"}
		}
	}

	return nil
}
