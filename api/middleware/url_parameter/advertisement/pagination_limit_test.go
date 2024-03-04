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

type LimitTestCase struct {
	testName string
	intValue int
	strValue string
}

var (
	rightParseLimitCases = []LimitTestCase{
		{
			testName: "Correct case",
			intValue: 10,
			strValue: "10",
		},
	}
	errorParseLimitCases = []LimitTestCase{
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

	rightValidateLimitCases = []LimitTestCase{
		{
			testName: "Correct case with value min limit (ege case)",
			intValue: utils.GetMinLimit(),
			strValue: strconv.Itoa(utils.GetMinLimit()),
		},
		{
			testName: "Correct case with value 24",
			intValue: 24,
			strValue: "24",
		},
		{
			testName: "Correct case with value max limit (ege case)",
			intValue: utils.GetMaxLimit(),
			strValue: strconv.Itoa(utils.GetMaxLimit()),
		},
	}
	errorValidateLimitCases = []LimitTestCase{
		{
			testName: "Error case with value negative",
			intValue: -24,
			strValue: "-24",
		},
		{
			testName: "Error case with value lower than min limit",
			intValue: utils.GetMinLimit() - 10,
			strValue: strconv.Itoa(utils.GetMinLimit() - 10),
		},
		{
			testName: "Error case with value higher than max limit",
			intValue: utils.GetMaxLimit() + 10,
			strValue: strconv.Itoa(utils.GetMaxLimit() + 10),
		},
	}

	rightLimitCases = []LimitTestCase{
		{testName: "Correct case not supply limit parameter"},
	}
	errorLimitCases = []LimitTestCase{
		{testName: "Error case supply too much limit parameter", strValue: "10"},
	}
)

func TestParseLimit(t *testing.T) {
	for _, limitCase := range rightParseLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			if val, err := parseLimit(limitCase.strValue); err != nil && val != limitCase.intValue {
				t.Errorf("parseLimit() should be sucess")
			}
		})
	}
	for _, limitCase := range errorParseLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			if _, err := parseLimit(limitCase.strValue); err == nil {
				t.Errorf("parseLimit() should be fail")
			}
		})
	}
}

func TestValidateLimit(t *testing.T) {
	for _, limitCase := range rightValidateLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			if status := validateLimit(limitCase.intValue); status == false {
				t.Errorf("validateLimit() should be sucess")
			}
		})
	}
	for _, limitCase := range errorValidateLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			if status := validateLimit(limitCase.intValue); status == true {
				t.Errorf("validateLimit() should not be sucess")
			}
		})
	}
}

func TestMiddlewarePaginationLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)

	for _, limitCase := range rightParseLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationLimit())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?limit="+limitCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, limitCase := range rightValidateLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationLimit())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?limit="+limitCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, limitCase := range rightLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationLimit())
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

	for _, limitCase := range errorParseLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationLimit())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?limit="+limitCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, limitCase := range errorValidateLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationLimit())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?limit="+limitCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, limitCase := range errorLimitCases {
		t.Run(limitCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewarePaginationLimit())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?limit="+limitCase.strValue+"&limit="+limitCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}
