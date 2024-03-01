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
	// make the string of age into int and error
	return strconv.Atoi(age)
}

// validate age parameter
func validateAge(age int) bool {
	// check the age is between minimum and maximum of age's limit
	return age >= utils.GetMinAge() && age <= utils.GetMaxAge()
}

// MiddlewareAge is to validate and parse the age parameter in url
func MiddlewareAge() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all age in url parameters
		ages, exists := c.Request.URL.Query()["age"]

		// set variable of the age
		var age int

		if !exists {
			// if the parameter of age not exist then make the age be default value
			age = utils.GetDefaultAge()
		} else if len(ages) == 0 {
			// for the case parameter of age supply, but it is empty.
			// for example, http://{HOST}/api/v1/ad?age=&platform=ios
			// then the ages will be [] (error case)
			e := failure.ClientError{Reason: "age's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(ages) == 1 {
			// parse the age parameter if parse error, it will return error's reason
			ageInt, err := parseAge(ages[0])
			age = ageInt
			if err != nil {
				e := failure.ClientError{Reason: "age's parameter parse error, " + err.Error()}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
			// validate the age parameter
			// to check if the value higher than or equal to minimum value and lower than or equal to maximum value
			// if validate fail, it will return error
			if status := validateAge(age); status == false {
				e := failure.ClientError{Reason: "age's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			// for the case parameter is too much
			// for example, http://{HOST}/api/v1/ad?age=10&age=5
			// then the ages will be ["10","5"] (error case)
			e := failure.ClientError{Reason: "age's parameter parse error, too much age parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		// set the age into context and continue to next handler
		c.Set("age", age)
		c.Next()
	}
}
