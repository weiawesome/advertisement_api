/*
The middleware of gin to validate and parse parameter in url about limit
*/

package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// parse limit parameter
func parseLimit(limit string) (int, error) {
	// make the string of limit into int and error
	return strconv.Atoi(limit)
}

// validate limit parameter
func validateLimit(limit int) bool {
	// check the limit is between minimum and maximum of limit
	return limit >= utils.GetMinLimit() && limit <= utils.GetMaxLimit()
}

// MiddlewarePaginationLimit is to validate and parse the limit parameter in url
func MiddlewarePaginationLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all limit in url parameters
		limits, exists := c.Request.URL.Query()["limit"]

		// set variable of the limit
		var limit int

		if !exists {
			// if the parameter of limit not exist then make the limit be default value
			limit = utils.GetDefaultLimit()
		} else if len(limits) == 0 {
			// for the case parameter of limit supply, but it is empty.
			// for example, http://{HOST}/api/v1/ad?limit=&platform=ios
			// then the limits will be [] (error case)
			e := failure.ClientError{Reason: "limit's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(limits) == 1 {
			// parse the limit parameter if parse error, it will return error's reason
			limitInt, err := parseLimit(limits[0])
			limit = limitInt
			if err != nil {
				e := failure.ClientError{Reason: "limit's parameter error, " + err.Error()}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
			// validate the limit parameter
			// to check if the value higher than or equal to minimum value
			// if validate fail, it will return error
			if status := validateLimit(limit); status == false {
				e := failure.ClientError{Reason: "limit's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			// for the case parameter is too much
			// for example, http://{HOST}/api/v1/ad?limit=10&limit=5
			// then the limits will be ["10","5"] (error case)
			e := failure.ClientError{Reason: "limit's parameter validate error, too much limit parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		// set the limit into context and continue to next handler
		c.Set("limit", limit)
		c.Next()
	}
}
