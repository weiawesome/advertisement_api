/*
The middleware of gin to validate and parse parameter in url about country
*/

package advertisement

import (
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CountryTestCase struct {
	testName   string
	countryISO string
}

var (
	rightValidateCountryCases = []CountryTestCase{
		{testName: "Correct case with TW", countryISO: "TW"},
		{testName: "Correct case with JP", countryISO: "JP"},
	}
	errorValidateCountryCases = []CountryTestCase{
		{testName: "Error case with Taiwan", countryISO: "Taiwan"},
		{testName: "Error case with Japan", countryISO: "Japan"},
	}

	rightCountryCases = []CountryTestCase{
		{testName: "Correct case not supply country parameter"},
	}
	errorCountryCases = []CountryTestCase{
		{testName: "Error case supply too much age parameter", countryISO: "TW"},
	}
)

// validate country parameter
func TestValidateCountry(t *testing.T) {
	utils.InitMaps()
	for _, countryCase := range rightValidateCountryCases {
		t.Run(countryCase.testName, func(t *testing.T) {
			if status := validateCountry(countryCase.countryISO); status == false {
				t.Errorf("validateCOuntry() should be sucess")
			}
		})
	}
	for _, countryCase := range errorValidateCountryCases {
		t.Run(countryCase.testName, func(t *testing.T) {
			if status := validateCountry(countryCase.countryISO); status == true {
				t.Errorf("validateCountry() should not be sucess")
			}
		})
	}
}

// MiddlewareCountry is to validate and parse the country parameter in url
func TestMiddlewareCountry(t *testing.T) {
	utils.InitMaps()
	for _, countryCase := range rightValidateCountryCases {
		t.Run(countryCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareCountry())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?country="+countryCase.countryISO, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, countryCase := range rightCountryCases {
		t.Run(countryCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareAge())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test", nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, countryCase := range errorValidateCountryCases {
		t.Run(countryCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareCountry())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?country="+countryCase.countryISO, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, countryCase := range errorCountryCases {
		t.Run(countryCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareCountry())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?country="+countryCase.countryISO+"&country="+countryCase.countryISO, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}

}
