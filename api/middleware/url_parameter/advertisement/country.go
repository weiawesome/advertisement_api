/*
The middleware of gin to validate and parse parameter in url about country
*/

package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// validate country parameter
func validateCountry(country string) bool {
	// check the country is in the country specified
	// countriesMap is make specified countries set into map[string]bool
	// if country in the specified countries set, the map of country value will be true
	return utils.GetCountriesMap()[country]
}

// MiddlewareCountry is to validate and parse the country parameter in url
func MiddlewareCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all country in url parameters
		countries, exists := c.Request.URL.Query()["country"]

		// set variable of the country
		var country string

		if !exists {
			// if the parameter of country not exist then make the country be default value
			// furthermore, it will view the query as no-constraint for the country
			country = utils.GetDefaultCountry()
		} else if len(countries) == 0 {
			// for the case parameter of country supply, but it is empty.
			// for example, http://{HOST}/api/v1/ad?country=&offset=3
			// then the countries will be [] (error case)
			e := failure.ClientError{Reason: "country's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(countries) == 1 {
			// validate the country parameter
			// to check if the value is in the specified country or not
			// if validate fail, it will return error
			country = countries[0]
			if status := validateCountry(country); status == false {
				e := failure.ClientError{Reason: "country's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			// for the case parameter is too much
			// for example, http://{HOST}/api/v1/ad?country=TW&country=JP
			// then the countries will be ["TW","JP"] (error case)
			e := failure.ClientError{Reason: "country's parameter validate error, too much country parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}
		// set the country into context and continue to next handler
		c.Set("country", country)
		c.Next()
	}
}
