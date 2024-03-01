/*
The middleware of gin to validate and parse parameter in url about age
*/

package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// parse age parameter
func parseAge(age string) (int, error) {
	return strconv.Atoi(age)
}

// validate age parameter
func validateAge(age int) bool {
	return age >= utils.GetMinAge() && age <= utils.GetMaxAge()
}

func MiddlewareAge() gin.HandlerFunc {
	return func(c *gin.Context) {
		ages, exists := c.Request.URL.Query()["age"]

		var age int

		if !exists {
			age = utils.GetDefaultAge()
		} else if len(ages) == 0 {
			e := failure.ClientError{Reason: "age's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(ages) == 1 {
			ageInt, err := parseAge(ages[0])
			age = ageInt
			if err != nil {
				e := failure.ClientError{Reason: "age's parameter parse error, " + err.Error()}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
			if status := validateAge(age); status == false {
				e := failure.ClientError{Reason: "age's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			e := failure.ClientError{Reason: "age's parameter parse error, too much age parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("age", age)
		c.Next()
	}
}
