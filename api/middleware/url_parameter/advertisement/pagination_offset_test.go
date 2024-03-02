/*
The middleware of gin to validate and parse parameter in url about offset
*/

package advertisement

import (
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type OffsetTestCase struct {
	testName string
	intValue int
	strValue string
}

var (
	rightParseOffsetCases = []OffsetTestCase{
		{
			testName: "Correct case",
			intValue: 10,
			strValue: "10",
		},
	}
	errorParseOffsetCases = []OffsetTestCase{
		{
			testName: "Error case with value not integer but a positive float",
			strValue: "9.9",
		},
		{
			testName: "Error case with value not integer but a negative float",
			strValue: "-9.9",
		},
		{
			testName: "Error case with value not integer but a string with char",
			strValue: "Error case",
		},
		{
			testName: "Error case with value empty",
			strValue: "",
		},
	}

	rightValidateOffsetCases = []OffsetTestCase{
		{
			testName: "Correct case with value min offset (ege case)",
			intValue: utils.GetMinOffset(),
			strValue: strconv.Itoa(utils.GetMinOffset()),
		},
		{
			testName: "Correct case with value 24",
			intValue: 24,
			strValue: "24",
		},
	}
	errorValidateOffsetCases = []OffsetTestCase{
		{
			testName: "Error case with value negative",
			intValue: -24,
			strValue: "-24",
		},
		{
			testName: "Error case with value lower than min offset",
			intValue: utils.GetMinOffset() - 10,
			strValue: strconv.Itoa(utils.GetMinOffset() - 10),
		},
	}

	rightOffsetCases = []OffsetTestCase{
		{testName: "Correct case not supply offset parameter"},
	}
	errorOffsetCases = []OffsetTestCase{
		{testName: "Error case supply too much offset parameter", strValue: "10"},
	}
)

// parse offset parameter
func TestParseOffset(t *testing.T) {
	for _, offsetCase := range rightParseOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			if val, err := parseOffset(offsetCase.strValue); err != nil && val != offsetCase.intValue {
				t.Errorf("parseOffset() should be sucess")
			}
		})
	}
	for _, offsetCase := range errorParseOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			if _, err := parseOffset(offsetCase.strValue); err == nil {
				t.Errorf("parseOffset() should be fail")
			}
		})
	}
}

// validate offset parameter
func TestValidateOffset(t *testing.T) {
	for _, offsetCase := range rightValidateOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			if status := validateOffset(offsetCase.intValue); status == false {
				t.Errorf("validateOffset() should be sucess")
			}
		})
	}
	for _, offsetCase := range errorValidateOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			if status := validateOffset(offsetCase.intValue); status == true {
				t.Errorf("validateOffset() should not be sucess")
			}
		})
	}
}

// MiddlewarePaginationOffset is to validate and parse the offset parameter in url
func TestMiddlewarePaginationOffset(t *testing.T) {
	gin.SetMode(gin.TestMode)

	for _, offsetCase := range rightParseOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationOffset())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?offset="+offsetCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, offsetCase := range rightValidateOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationOffset())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?offset="+offsetCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, offsetCase := range rightOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationOffset())
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

	for _, offsetCase := range errorParseOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationOffset())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?offset="+offsetCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, offsetCase := range errorValidateOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationOffset())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?offset="+offsetCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, offsetCase := range errorOffsetCases {
		t.Run(offsetCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationOffset())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?offset="+offsetCase.strValue+"&offset="+offsetCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}
