package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func validateAge(age int) bool {
	if age != utils.GetDefaultAge() && (age < utils.GetMinAge() || age > utils.GetMaxAge()) {
		return false
	}
	return true
}

func parseAge(age string) (int, error) {
	if age == "" {
		return utils.GetDefaultAge(), nil
	} else {
		return strconv.Atoi(age)
	}
}

func MiddlewareAge() gin.HandlerFunc {
	return func(c *gin.Context) {
		age := c.DefaultQuery("age", "")

		ageInt, err := parseAge(age)
		if err != nil {
			e := failure.ClientError{Reason: "age's parameter parse error, " + err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		if status := validateAge(ageInt); status == false {
			e := failure.ClientError{Reason: "age's parameter validate error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("age", ageInt)
		c.Next()
	}
}
