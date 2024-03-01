/*
The middleware of gin to validate and parse parameter in url about offset
*/

package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// parse offset parameter
func parseOffset(offset string) (int, error) {
	// make the string of offset into int and error
	return strconv.Atoi(offset)
}

// validate offset parameter
func validateOffset(offset int) bool {
	// check the offset is larger than or equal to age's minimum
	return offset >= utils.GetMinOffset()
}

// MiddlewarePaginationOffset is to validate and parse the offset parameter in url
func MiddlewarePaginationOffset() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all offset in url parameters
		offsets, exists := c.Request.URL.Query()["offset"]

		// set variable of the offset
		var offset int

		if !exists {
			// if the parameter of offset not exist then make the age be default value
			offset = utils.GetDefaultOffset()
		} else if len(offsets) == 0 {
			// for the case parameter of offset supply, but it is empty.
			// for example, http://{HOST}/api/v1/ad?offset=&age=5
			// then the offsets will be [] (error case)
			e := failure.ClientError{Reason: "offset's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(offsets) == 1 {
			// parse the offset parameter if parse error, it will return error's reason
			offsetInt, err := parseOffset(offsets[0])
			offset = offsetInt
			if err != nil {
				e := failure.ClientError{Reason: "offset's parameter error, " + err.Error()}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
			// validate the age parameter
			// to check if the value higher than or equal to minimum value
			// if validate fail, it will return error
			if status := validateOffset(offset); status == false {
				e := failure.ClientError{Reason: "offset's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			// for the case parameter is too much
			// for example, http://{HOST}/api/v1/ad?offset=10&offset=5
			// then the offsets will be ["10","5"] (error case)
			e := failure.ClientError{Reason: "offset's parameter validate error, too much offset parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		// set the offset into context and continue to next handler
		c.Set("offset", offset)
		c.Next()
	}
}
