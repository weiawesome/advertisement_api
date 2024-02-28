package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateGender(gender string) bool {
	return gender == utils.GetDefaultGender() || utils.GetGendersMap()[gender]
}

func MiddlewareGender() gin.HandlerFunc {
	return func(c *gin.Context) {
		gender := c.DefaultQuery("gender", utils.GetDefaultGender())

		if status := validateGender(gender); status == false {
			e := failure.ClientError{Reason: "gender's parameter validate error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("gender", gender)
		c.Next()
	}
}
