/*
The middleware of gin to validate and parse parameter in url about platform
*/

package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// validate platform parameter
func validatePlatform(platform string) bool {
	// check the platform is in the platform specified
	// platformsMap is make specified platforms set into map[string]bool
	// if platform in the specified platforms set, the map of platform value will be true
	return utils.GetPlatformsMap()[platform]
}

// MiddlewarePlatform is to validate and parse the platform parameter in url
func MiddlewarePlatform() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all platform in url parameters
		platforms, exists := c.Request.URL.Query()["platform"]

		// set variable of the platform
		var platform string

		if !exists {
			// if the parameter of platform not exist then make the platform be default value
			// furthermore, it will view the query as no-constraint for the platform
			platform = utils.GetDefaultPlatform()
		} else if len(platforms) == 0 {
			// for the case parameter of platform supply, but it is empty.
			// for example, http://{HOST}/api/v1/ad?platform=&gender=M
			// then the platforms will be [] (error case)
			e := failure.ClientError{Reason: "platform's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(platforms) == 1 {
			// validate the platform parameter
			// to check if the value is in the specified platforms or not
			// if validate fail, it will return error
			platform = platforms[0]
			if status := validatePlatform(platform); status == false {
				e := failure.ClientError{Reason: "platform's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			// for the case parameter is too much
			// for example, http://{HOST}/api/v1/ad?platform=ios&platform=android
			// then the platforms will be ["ios","android"] (error case)
			e := failure.ClientError{Reason: "platform's parameter validate error, too much platform parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		// set the country into context and continue to next handler
		c.Set("platform", platform)
		c.Next()
	}
}
