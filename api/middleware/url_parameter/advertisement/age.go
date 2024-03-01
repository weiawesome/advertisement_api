/*
The middleware of gin to validate and parse parameter in url about age
*/

package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// parse and validate the age parameter
func parseValidateAge(age string) (int, error) {
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		return ageInt, err
	}

	if ageInt < utils.GetMinAge() || ageInt > utils.GetMaxAge() {
		return ageInt, errors.New("invalidate age")
	}

	return ageInt, nil
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
			ageInt, err := parseValidateAge(ages[0])
			age = ageInt
			if err != nil {
				e := failure.ClientError{Reason: "age's parameter parse error, " + err.Error()}
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
