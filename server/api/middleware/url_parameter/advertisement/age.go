package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func parseValidateAge(age string) (int, error) {
	if age == "" {
		return utils.GetDefaultAge(), nil
	} else {

		age, err := strconv.Atoi(age)
		if err != nil {
			return age, err
		}

		if age < utils.GetMinAge() || age > utils.GetMaxAge() {
			return age, errors.New("invalidate age")
		} else {
			return age, err
		}
	}
}

func MiddlewareAge() gin.HandlerFunc {
	return func(c *gin.Context) {
		age := c.DefaultQuery("age", "")

		ageInt, err := parseValidateAge(age)
		if err != nil {
			e := failure.ClientError{Reason: "age's parameter parse error, " + err.Error()}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("age", ageInt)
		c.Next()
	}
}
