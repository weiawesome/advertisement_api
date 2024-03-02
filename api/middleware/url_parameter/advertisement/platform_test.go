/*
The middleware of gin to validate and parse parameter in url about platform
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

type PlatformTestCase struct {
	testName     string
	platformCode string
}

var (
	rightValidatePlatformCases = []PlatformTestCase{
		{testName: "Correct case with ios", platformCode: "ios"},
		{testName: "Correct case with android", platformCode: "android"},
		{testName: "Correct case with web", platformCode: "web"},
	}
	errorValidatePlatformCases = []PlatformTestCase{
		{testName: "Error case with switch", platformCode: "switch"},
		{testName: "Error case with ps5", platformCode: "ps5"},
		{testName: "Error case with steam", platformCode: "steam"},
	}

	rightPlatformCases = []PlatformTestCase{
		{testName: "Correct case not supply platform parameter"},
	}
	errorPlatformCases = []PlatformTestCase{
		{testName: "Error case supply too much age parameter", platformCode: "ios"},
	}
)

// validate platform parameter
func TestValidatePlatform(t *testing.T) {
	utils.InitMaps()
	for _, platformCase := range rightValidatePlatformCases {
		t.Run(platformCase.testName, func(t *testing.T) {
			if status := validatePlatform(platformCase.platformCode); status == false {
				t.Errorf("validateCOuntry() should be sucess")
			}
		})
	}
	for _, platformCase := range errorValidatePlatformCases {
		t.Run(platformCase.testName, func(t *testing.T) {
			if status := validatePlatform(platformCase.platformCode); status == true {
				t.Errorf("validatePlatform() should not be sucess")
			}
		})
	}
}

// MiddlewarePlatform is to validate and parse the platform parameter in url
func TestMiddlewarePlatform(t *testing.T) {
	utils.InitMaps()
	for _, platformCase := range rightValidatePlatformCases {
		t.Run(platformCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePlatform())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?platform="+platformCase.platformCode, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, platformCase := range rightPlatformCases {
		t.Run(platformCase.testName, func(t *testing.T) {
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
	for _, platformCase := range errorValidatePlatformCases {
		t.Run(platformCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePlatform())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?platform="+platformCase.platformCode, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, platformCase := range errorPlatformCases {
		t.Run(platformCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePlatform())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?platform="+platformCase.platformCode+"&platform="+platformCase.platformCode, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}
