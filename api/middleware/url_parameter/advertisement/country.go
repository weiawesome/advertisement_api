package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validateCountry(country string) bool {
	return country == utils.GetDefaultCountry() || utils.GetCountriesMap()[country]
}

func MiddlewareCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		country := c.DefaultQuery("country", utils.GetDefaultCountry())

		if status := validateCountry(country); status == false {
			e := failure.ClientError{Reason: "country's parameter validate error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		c.Set("country", country)
		c.Next()
	}
}
