package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func validateLimit(limit int) bool {
	return limit >= utils.GetMinLimit() && limit <= utils.GetMaxLimit()
}
func parseLimit(limit string) (int, error) {
	return strconv.Atoi(limit)
}

func MiddlewarePaginationLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		limits, exists := c.Request.URL.Query()["limit"]

		var limit int

		if !exists {
			limit = utils.GetDefaultLimit()
		} else if len(limits) == 0 {
			e := failure.ClientError{Reason: "limit's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(limits) == 1 {
			limitInt, err := parseLimit(limits[0])
			limit = limitInt
			if err != nil {
				e := failure.ClientError{Reason: "limit's parameter error, " + err.Error()}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
			if status := validateLimit(limit); status == false {
				e := failure.ClientError{Reason: "limit's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			e := failure.ClientError{Reason: "limit's parameter validate error, too much limit parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("limit", limit)
		c.Next()
	}
}
