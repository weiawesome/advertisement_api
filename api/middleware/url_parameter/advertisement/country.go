package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateCountry(country string) bool {
	return utils.GetCountriesMap()[country]
}

func MiddlewareCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		countries, exists := c.Request.URL.Query()["country"]

		var country string

		if !exists {
			country = utils.GetDefaultCountry()
		} else if len(countries) == 0 {
			e := failure.ClientError{Reason: "country's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(countries) == 1 {
			country = countries[0]
			if status := validateCountry(country); status == false {
				e := failure.ClientError{Reason: "country's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			e := failure.ClientError{Reason: "country's parameter validate error, too much country parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("country", country)
		c.Next()
	}
}
