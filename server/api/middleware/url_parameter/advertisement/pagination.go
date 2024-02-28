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

func parseOffset(offset string) (int, error) {
	if offset == "" {
		return utils.GetDefaultOffset(), nil
	} else {
		return strconv.Atoi(offset)
	}
}
func parseLimit(limit string) (int, error) {
	if limit == "" {
		return utils.GetDefaultLimit(), nil
	} else {
		return strconv.Atoi(limit)
	}
}

func MiddlewarePagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		offset := c.DefaultQuery("offset", "")
		limit := c.DefaultQuery("limit", "")

		offsetInt, offsetErr := parseOffset(offset)
		if offsetErr != nil {
			e := failure.ClientError{Reason: "offset's parameter error, " + offsetErr.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		limitInt, limitErr := parseLimit(limit)
		if limitErr != nil {
			e := failure.ClientError{Reason: "offset's parameter error, " + limitErr.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		if status := validateLimit(limitInt); status == false {
			e := failure.ClientError{Reason: "limit's parameter validate error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("offset", offsetInt)
		c.Set("limit", limitInt)
		c.Next()
	}
}
