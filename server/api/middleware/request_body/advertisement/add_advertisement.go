package advertisement

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func isSubSet(childSet []string, parentMap map[string]bool) bool {
	var result, val bool
	for _, child := range childSet {
		val, result = parentMap[child]
		if (!val) || (!result) {
			return result
		}
		parentMap[child] = false
	}
	return result
}
func validateCondition(data advertisement.AddAdvertisementRequest) error {
	if data.Conditions.AgeStart == nil && data.Conditions.AgeEnd == nil {
		return nil
	} else if data.Conditions.AgeStart == nil || data.Conditions.AgeEnd == nil {
		return errors.New("age limit must be a pair")
	} else if *data.Conditions.AgeStart > *data.Conditions.AgeEnd {
		return errors.New("age limit is illegal")
	} else if data.Conditions.Gender != nil && !isSubSet(data.Conditions.Gender, utils.GetGendersMap()) {
		return errors.New("invalid gender")
	} else if data.Conditions.Platform != nil && !isSubSet(data.Conditions.Platform, utils.GetPlatformsMap()) {
		return errors.New("invalid platform")
	} else if data.Conditions.Country != nil && !isSubSet(data.Conditions.Country, utils.GetCountriesMap()) {
		return errors.New("invalid country")
	}
	return nil
}

func validateBasicInfo(data advertisement.AddAdvertisementRequest) error {
	if data.StartAt.IsZero() || data.EndAt.IsZero() {
		return errors.New("startAt and endAt parameter is required")
	} else if data.EndAt.Before(data.StartAt) || data.EndAt.Equal(data.StartAt) {
		return errors.New("endAt can't be before or equal to startAt")
	} else if data.Title == nil {
		return errors.New("title is required")
	}

	return nil
}

func MiddlewareAddAdvertisement() gin.HandlerFunc {
	return func(c *gin.Context) {

		var data advertisement.AddAdvertisementRequest

		if err := c.ShouldBindJSON(&data); err != nil {
			e := failure.ClientError{Reason: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		err := validateBasicInfo(data)
		if err != nil {
			e := failure.ClientError{Reason: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		err = validateCondition(data)
		if err != nil {
			e := failure.ClientError{Reason: err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("data", data)
		c.Next()
	}
}
