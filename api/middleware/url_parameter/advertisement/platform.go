package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validatePlatform(platform string) bool {
	return platform == utils.GetDefaultPlatform() || utils.GetPlatformsMap()[platform]
}

func MiddlewarePlatform() gin.HandlerFunc {
	return func(c *gin.Context) {
		platform := c.DefaultQuery("platform", utils.GetDefaultPlatform())

		if status := validatePlatform(platform); status == false {
			e := failure.ClientError{Reason: "platform's parameter validate error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("platform", platform)
		c.Next()
	}
}
