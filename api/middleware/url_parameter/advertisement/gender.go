package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateGender(gender string) bool {
	return utils.GetGendersMap()[gender]
}

func MiddlewareGender() gin.HandlerFunc {
	return func(c *gin.Context) {
		genders, exists := c.Request.URL.Query()["gender"]

		var gender string

		if !exists {
			gender = utils.GetDefaultGender()
		} else if len(genders) == 0 {
			e := failure.ClientError{Reason: "gender's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(genders) == 1 {
			gender = genders[0]
			if status := validateGender(gender); status == false {
				e := failure.ClientError{Reason: "gender's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			e := failure.ClientError{Reason: "gender's parameter validate error, too much gender parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("gender", gender)
		c.Next()
	}
}
