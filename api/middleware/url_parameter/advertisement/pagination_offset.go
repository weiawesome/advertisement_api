package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func parseOffset(offset string) (int, error) {
	return strconv.Atoi(offset)
}
func validateOffset(offset int) bool {
	return offset >= utils.GetMinOffset()
}

func MiddlewarePaginationOffset() gin.HandlerFunc {
	return func(c *gin.Context) {
		offsets, exists := c.Request.URL.Query()["offset"]

		var offset int

		if !exists {
			offset = utils.GetDefaultOffset()
		} else if len(offsets) == 0 {
			e := failure.ClientError{Reason: "offset's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(offsets) == 1 {
			offsetInt, err := parseOffset(offsets[0])
			offset = offsetInt
			if err != nil {
				e := failure.ClientError{Reason: "offset's parameter error, " + err.Error()}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
			if status := validateOffset(offset); status == false {
				e := failure.ClientError{Reason: "offset's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			e := failure.ClientError{Reason: "offset's parameter validate error, too much offset parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("offset", offset)
		c.Next()
	}
}
