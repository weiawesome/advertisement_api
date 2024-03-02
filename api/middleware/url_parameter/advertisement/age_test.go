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

type AgeTestCase struct {
	testName string
	intValue int
	strValue string
}

var (
	rightParseAgeCases = []AgeTestCase{
		{
			testName: "Correct case",
			intValue: 10,
			strValue: "10",
		},
	}
	errorParseAgeCases = []AgeTestCase{
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

	rightValidateAgeCases = []AgeTestCase{
		{
			testName: "Correct case with value min age (ege case)",
			intValue: utils.GetMinAge(),
			strValue: strconv.Itoa(utils.GetMinAge()),
		},
		{
			testName: "Correct case with value 24",
			intValue: 24,
			strValue: "24",
		},
		{
			testName: "Correct case with value max age (ege case)",
			intValue: utils.GetMaxAge(),
			strValue: strconv.Itoa(utils.GetMaxAge()),
		},
	}
	errorValidateAgeCases = []AgeTestCase{
		{
			testName: "Error case with value negative",
			intValue: -24,
			strValue: "-24",
		},
		{
			testName: "Error case with value lower than min age",
			intValue: utils.GetMinAge() - 10,
			strValue: strconv.Itoa(utils.GetMinAge() - 10),
		},
		{
			testName: "Error case with value higher than max age",
			intValue: utils.GetMaxAge() + 10,
			strValue: strconv.Itoa(utils.GetMaxAge() + 10),
		},
	}

	rightAgeCases = []AgeTestCase{
		{testName: "Correct case not supply age parameter"},
	}
	errorAgeCases = []AgeTestCase{
		{testName: "Error case supply too much age parameter", strValue: "10"},
	}
)

func TestParseAge(t *testing.T) {
	for _, ageCase := range rightParseAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			if val, err := parseAge(ageCase.strValue); err != nil && val != ageCase.intValue {
				t.Errorf("parseAge() should be sucess")
			}
		})
	}
	for _, ageCase := range errorParseAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			if _, err := parseAge(ageCase.strValue); err == nil {
				t.Errorf("parseAge() should be fail")
			}
		})
	}
}

func TestValidateAge(t *testing.T) {
	for _, ageCase := range rightValidateAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			if status := validateAge(ageCase.intValue); status == false {
				t.Errorf("validateAge() should be sucess")
			}
		})
	}
	for _, ageCase := range errorValidateAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			if status := validateAge(ageCase.intValue); status == true {
				t.Errorf("validateAge() should not be sucess")
			}
		})
	}
}

func TestMiddlewareAge(t *testing.T) {
	gin.SetMode(gin.TestMode)

	for _, ageCase := range rightParseAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareAge())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?age="+ageCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, ageCase := range rightValidateAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareAge())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?age="+ageCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Passed", w.Body.String())
		})
	}
	for _, ageCase := range rightAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
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

	for _, ageCase := range errorParseAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareAge())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?age="+ageCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, ageCase := range errorValidateAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareAge())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?age="+ageCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
	for _, ageCase := range errorAgeCases {
		t.Run(ageCase.testName, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareAge())
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/test?age="+ageCase.strValue+"&age="+ageCase.strValue, nil)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}
