/*
The middleware of gin to validate and parse parameter in url about gender
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

type GenderTestCase struct {
	testName   string
	genderCode string
}

var (
	rightValidateGenderCases = []GenderTestCase{
		{testName: "Correct case with M", genderCode: "F"},
		{testName: "Correct case with F", genderCode: "M"},
	}
	errorValidateGenderCases = []GenderTestCase{
		{testName: "Error case with T", genderCode: "T"},
	}

	rightGenderCases = []GenderTestCase{
		{testName: "Correct case not supply gender parameter"},
	}
	errorGenderCases = []GenderTestCase{
		{testName: "Error case supply too much age parameter", genderCode: "M"},
	}
)

// validate gender parameter
func TestValidateGender(t *testing.T) {
	utils.InitMaps()
	for _, genderCase := range rightValidateGenderCases {
		t.Run(genderCase.testName, func(t *testing.T) {
			if status := validateGender(genderCase.genderCode); status == false {
				t.Errorf("validateCOuntry() should be sucess")
			}
		})
	}
	for _, genderCase := range errorValidateGenderCases {
		t.Run(genderCase.testName, func(t *testing.T) {
			if status := validateGender(genderCase.genderCode); status == true {
				t.Errorf("validateGender() should not be sucess")
			}
		})
	}
}

// MiddlewareGender is to validate and parse the gender parameter in url
func TestMiddlewareGender(t *testing.T) {
	utils.InitMaps()
	for _, genderCase := range rightValidateGenderCases {
		t.Run(genderCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareGender())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?gender="+genderCase.genderCode, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, genderCase := range rightGenderCases {
		t.Run(genderCase.testName, func(t *testing.T) {
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
	for _, genderCase := range errorValidateGenderCases {
		t.Run(genderCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareGender())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?gender="+genderCase.genderCode, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, genderCase := range errorGenderCases {
		t.Run(genderCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareGender())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?gender="+genderCase.genderCode+"&gender="+genderCase.genderCode, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}
