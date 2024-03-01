package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validatePlatform(platform string) bool {
	return utils.GetPlatformsMap()[platform]
}

func MiddlewarePlatform() gin.HandlerFunc {
	return func(c *gin.Context) {
		platforms, exists := c.Request.URL.Query()["platform"]

		var platform string

		if !exists {
			platform = utils.GetDefaultPlatform()
		} else if len(platforms) == 0 {
			e := failure.ClientError{Reason: "platform's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(platforms) == 1 {
			platform = platforms[0]
			if status := validatePlatform(platform); status == false {
				e := failure.ClientError{Reason: "platform's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			e := failure.ClientError{Reason: "platform's parameter validate error, too much platform parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("platform", platform)
		c.Next()
	}
}
