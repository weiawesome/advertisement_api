/*
The middleware to check request body's content is regular or not
It will check all the content include genders,country, age, time etc.
*/

package advertisement

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// check the element is subset in specified set
// It let the specified set into map and let the set's value in map be true.
// for example, genders = { "M" , "F" } then GendersMap = { "M" : true, "F" : true }
// hence, it can easily to check whether element is in the set or not.
func isSubSet(childSet []string, parentMap map[string]bool) bool {
	// if child set is nil (without supply) then return true
	if childSet == nil {
		return true
	}

	var result, val bool
	existMap := make(map[string]bool)

	// enumerate all the element check the value
	for _, child := range childSet {
		val, result = parentMap[child]

		// if it not exists in map or its value is false then return false
		if (!val) || (!result) {
			return false
		}

		// if the value has been appeared, then return false
		if exist, _ := existMap[child]; exist == true {
			return false
		}
		existMap[child] = true
	}

	// it is subset in the specified set then return true
	return true
}

// validate condition in the add advertisement request
func validateCondition(data advertisement.AddAdvertisementRequest) error {
	if (data.Conditions.AgeStart == nil) != (data.Conditions.AgeEnd == nil) {
		// AgeStart and AgeEnd just only supply one
		return errors.New("age limit must be a pair")
	} else if data.Conditions.AgeStart != nil && data.Conditions.AgeEnd != nil && (*data.Conditions.AgeStart > *data.Conditions.AgeEnd) {
		// AgeStart larger than AgeEnd
		return errors.New("age limit is illegal")
	} else if !isSubSet(data.Conditions.Gender, utils.GetGendersMap()) {
		// the gender is not in specified gender set
		return errors.New("invalid gender")
	} else if !isSubSet(data.Conditions.Platform, utils.GetPlatformsMap()) {
		// the platform is not in specified platform set
		return errors.New("invalid platform")
	} else if !isSubSet(data.Conditions.Country, utils.GetCountriesMap()) {
		// the country is not in specified country set
		return errors.New("invalid country")
	}
	// check all condition is ok
	return nil
}

// to validate basic information in the request (all information in here is required)
func validateBasicInfo(data advertisement.AddAdvertisementRequest) error {
	if data.StartAt == nil || data.EndAt == nil {
		// lack of StartAt or EndAt
		return errors.New("startAt and endAt parameter is required")
	} else if (*data.EndAt).Before(*data.StartAt) || (*data.EndAt).Equal(*data.StartAt) {
		// EndAt is before or equal to StartAt
		return errors.New("endAt can't be before or equal to startAt")
	} else if data.Title == nil {
		// lack of title
		return errors.New("title is required")
	}

	return nil
}

// MiddlewareAddAdvertisement is to validate add advertisement body information
func MiddlewareAddAdvertisement() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the data then bind it. if it fails, return error.
		var data advertisement.AddAdvertisementRequest
		if err := c.ShouldBindJSON(&data); err != nil {
			e := failure.ClientError{Reason: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		// validate basic information in request including time and title
		// if validation is fail, it will return error to the client
		err := validateBasicInfo(data)
		if err != nil {
			e := failure.ClientError{Reason: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		// validate condition information in request including genders, country, platform,time.
		// if validation is fail, it will return error to the client
		err = validateCondition(data)
		if err != nil {
			e := failure.ClientError{Reason: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		// set the data in context and continue
		c.Set("data", data)
		c.Next()
	}
}
