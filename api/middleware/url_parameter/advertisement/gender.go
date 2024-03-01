/*
The middleware of gin to validate and parse parameter in url about gender
*/

package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// validate gender parameter
func validateGender(gender string) bool {
	// check the gender is in the gender specified
	// gendersMap is make specified genders set into map[string]bool
	// if gender in the specified genders set, the map of gender value will be true
	return utils.GetGendersMap()[gender]
}

// MiddlewareGender is to validate and parse the gender parameter in url
func MiddlewareGender() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get all gender in url parameters
		genders, exists := c.Request.URL.Query()["gender"]

		// set variable of the gender
		var gender string

		if !exists {
			// if the parameter of gender not exist then make the gender be default value
			// furthermore, it will view the query as no-constraint for the gender
			gender = utils.GetDefaultGender()
		} else if len(genders) == 0 {
			// for the case parameter of gender supply, but it is empty.
			// for example, http://{HOST}/api/v1/ad?gender=&country=TW
			// then the genders will be [] (error case)
			e := failure.ClientError{Reason: "gender's parameter error, value empty error"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		} else if len(genders) == 1 {
			// validate the gender parameter
			// to check if the value is in the specified genders or not
			// if validate fail, it will return error
			gender = genders[0]
			if status := validateGender(gender); status == false {
				e := failure.ClientError{Reason: "gender's parameter validate error"}
				c.JSON(http.StatusBadRequest, e)
				c.Abort()
				return
			}
		} else {
			// for the case parameter is too much
			// for example, http://{HOST}/api/v1/ad?gender=M&gender=F
			// then the countries will be ["M","F"] (error case)
			e := failure.ClientError{Reason: "gender's parameter validate error, too much gender parameters"}
			c.JSON(http.StatusBadRequest, e)
			c.Abort()
			return
		}

		// set the gender into context and continue to next handler
		c.Set("gender", gender)
		c.Next()
	}
}
