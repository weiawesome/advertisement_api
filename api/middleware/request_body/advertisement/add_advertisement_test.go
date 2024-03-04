/*
The middleware to check request body's content is regular or not
It will check all the content include genders,country, age, time etc.
*/

package advertisement

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/utils"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type IsSubsetCase struct {
	testName  string
	childSet  []string
	parentMap map[string]bool
}
type AdvertisementCase struct {
	testName      string
	advertisement advertisement.AddAdvertisementRequest
}

var (
	rightIsSubSetCases = []IsSubsetCase{
		{
			testName:  "Correct case",
			childSet:  []string{"key-1", "key-2"},
			parentMap: map[string]bool{"key-1": true, "key-2": true, "key-3": true},
		},
		{
			testName:  "Correct case with nil child set",
			childSet:  nil,
			parentMap: map[string]bool{"key-1": true, "key-2": true, "key-3": true},
		},
	}
	errorIsSubSetCases = []IsSubsetCase{
		{
			testName:  "Error case with duplicate key in child set",
			childSet:  []string{"key-1", "key-2", "key-1"},
			parentMap: map[string]bool{"key-1": true, "key-2": true, "key-3": true},
		},
		{
			testName:  "Error case with not exist key in child set",
			childSet:  []string{"key-4"},
			parentMap: map[string]bool{"key-1": true, "key-2": true, "key-3": true},
		},
	}

	rightValidateConditionCases = []AdvertisementCase{
		{
			testName: "Correct case supply all condition",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(10),
					AgeEnd:   func(i int) *int { return &i }(30),
					Gender:   []string{"M", "F"},
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Correct case without supply age",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: nil,
					AgeEnd:   nil,
					Gender:   []string{"M", "F"},
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Correct case without supply gender",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(10),
					AgeEnd:   func(i int) *int { return &i }(30),
					Gender:   nil,
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Correct case without supply country",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(10),
					AgeEnd:   func(i int) *int { return &i }(30),
					Gender:   []string{"M", "F"},
					Country:  nil,
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Correct case without supply platform",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(10),
					AgeEnd:   func(i int) *int { return &i }(30),
					Gender:   []string{"M", "F"},
					Country:  []string{"TW", "JP"},
					Platform: nil,
				},
			},
		},
		{
			testName: "Correct case without supply anything",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: nil,
					AgeEnd:   nil,
					Gender:   nil,
					Country:  nil,
					Platform: nil,
				},
			},
		},
	}
	errorValidateConditionCases = []AdvertisementCase{
		{
			testName: "Error case supply only age start",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(10),
					AgeEnd:   nil,
					Gender:   []string{"M", "F"},
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Error case supply only age end",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: nil,
					AgeEnd:   func(i int) *int { return &i }(30),
					Gender:   []string{"M", "F"},
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Error case supply age logic error",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(30),
					AgeEnd:   func(i int) *int { return &i }(10),
					Gender:   []string{"M", "F"},
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Error case supply gender not exist",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: nil,
					AgeEnd:   nil,
					Gender:   []string{"T"},
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Error case supply country not exist",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: nil,
					AgeEnd:   nil,
					Gender:   nil,
					Country:  []string{"Taiwan"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Error case supply platform not exist",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: nil,
					AgeEnd:   nil,
					Gender:   nil,
					Country:  nil,
					Platform: []string{"switch"},
				},
			},
		},
		{
			testName: "Correct case supply gender duplicate",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(10),
					AgeEnd:   func(i int) *int { return &i }(30),
					Gender:   []string{"M", "M"},
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Correct case supply country duplicate",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(10),
					AgeEnd:   func(i int) *int { return &i }(30),
					Gender:   []string{"M", "F"},
					Country:  []string{"TW", "TW"},
					Platform: []string{"ios", "android"},
				},
			},
		},
		{
			testName: "Correct case supply platform duplicate",
			advertisement: advertisement.AddAdvertisementRequest{
				Conditions: advertisement.Condition{
					AgeStart: func(i int) *int { return &i }(10),
					AgeEnd:   func(i int) *int { return &i }(30),
					Gender:   []string{"M", "F"},
					Country:  []string{"TW", "JP"},
					Platform: []string{"ios", "ios"},
				},
			},
		},
	}

	rightValidateBasicInfoCases = []AdvertisementCase{
		{
			testName: "Correct case",
			advertisement: advertisement.AddAdvertisementRequest{
				Title:   func(s string) *string { return &s }("Correct case's title"),
				StartAt: func(t time.Time) *time.Time { return &t }(time.Now().UTC()),
				EndAt:   func(t time.Time) *time.Time { return &t }(time.Now().UTC().Add(time.Hour * 10)),
			},
		},
	}
	errorValidateBasicInfoCases = []AdvertisementCase{
		{
			testName: "Error case with not supply title",
			advertisement: advertisement.AddAdvertisementRequest{
				Title:   nil,
				StartAt: func(t time.Time) *time.Time { return &t }(time.Now().UTC()),
				EndAt:   func(t time.Time) *time.Time { return &t }(time.Now().UTC().Add(time.Hour * 10)),
			},
		},
		{
			testName: "Error case with not supply start at",
			advertisement: advertisement.AddAdvertisementRequest{
				Title:   func(s string) *string { return &s }("Correct case's title"),
				StartAt: nil,
				EndAt:   func(t time.Time) *time.Time { return &t }(time.Now().UTC().Add(time.Hour * 10)),
			},
		},
		{
			testName: "Error case with not supply end at",
			advertisement: advertisement.AddAdvertisementRequest{
				Title:   func(s string) *string { return &s }("Correct case's title"),
				StartAt: func(t time.Time) *time.Time { return &t }(time.Now().UTC()),
				EndAt:   nil,
			},
		},
		{
			testName: "Error case with time logic error",
			advertisement: advertisement.AddAdvertisementRequest{
				Title:   func(s string) *string { return &s }("Correct case's title"),
				StartAt: func(t time.Time) *time.Time { return &t }(time.Now().UTC().Add(time.Hour * 10)),
				EndAt:   func(t time.Time) *time.Time { return &t }(time.Now().UTC()),
			},
		},
	}
)

func TestIsSubSet(t *testing.T) {
	for _, rightCase := range rightIsSubSetCases {
		t.Run(rightCase.testName, func(t *testing.T) {
			if result := isSubSet(rightCase.childSet, rightCase.parentMap); result != true {
				t.Errorf("isSubSet() = %v, want %v", true, result)
			}
		})
	}
	for _, errorCase := range errorIsSubSetCases {
		t.Run(errorCase.testName, func(t *testing.T) {
			if result := isSubSet(errorCase.childSet, errorCase.parentMap); result != false {
				t.Errorf("isSubSet() = %v, want %v", false, result)
			}
		})
	}
}

func TestValidateCondition(t *testing.T) {
	utils.InitMaps()
	for _, rightCase := range rightValidateConditionCases {
		t.Run(rightCase.testName, func(t *testing.T) {
			if err := validateCondition(rightCase.advertisement); err != nil {
				t.Errorf("validateCondition() = %v, want %v", err.Error(), nil)
			}
		})
	}
	for _, errorCase := range errorValidateConditionCases {
		t.Run(errorCase.testName, func(t *testing.T) {
			if err := validateCondition(errorCase.advertisement); err == nil {
				t.Errorf("validateCondition() = %v, want %v", err.Error(), nil)
			}
		})
	}
}

func TestValidateBasicInfo(t *testing.T) {
	for _, rightCase := range rightValidateBasicInfoCases {
		t.Run(rightCase.testName, func(t *testing.T) {
			if err := validateBasicInfo(rightCase.advertisement); err != nil {
				t.Errorf("validateBasicInfo() = %v, want %v", err.Error(), nil)
			}
		})
	}
	for _, errorCase := range errorValidateBasicInfoCases {
		t.Run(errorCase.testName, func(t *testing.T) {
			if err := validateBasicInfo(errorCase.advertisement); err == nil {
				t.Errorf("validateBasicInfo() = %v, want %v", err.Error(), nil)
			}
		})
	}
}

func TestMiddlewareAddAdvertisement(t *testing.T) {
	utils.InitMaps()

	gin.SetMode(gin.TestMode)

	t.Run("Error case json not bind", func(t *testing.T) {
		router := gin.New()
		router.Use(MiddlewareAddAdvertisement())
		router.POST("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "Passed")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	for _, basicInfoCase := range rightValidateBasicInfoCases {
		for _, conditionCase := range rightValidateConditionCases {
			t.Run("With correct info and correct condition"+basicInfoCase.testName+conditionCase.testName, func(t *testing.T) {

				testCase := advertisement.AddAdvertisementRequest{
					Title:      basicInfoCase.advertisement.Title,
					StartAt:    basicInfoCase.advertisement.StartAt,
					EndAt:      basicInfoCase.advertisement.EndAt,
					Conditions: conditionCase.advertisement.Conditions,
				}

				jsonData, err := json.Marshal(testCase)
				if err != nil {
					t.Errorf("Error to jsonify data with error " + err.Error())
					return
				}

				router := gin.New()
				router.Use(MiddlewareAddAdvertisement())
				router.POST("/test", func(c *gin.Context) {
					c.String(http.StatusOK, "Passed")
				})

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))

				router.ServeHTTP(w, req)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, "Passed", w.Body.String())
			})
		}
	}

	for _, basicInfoCase := range rightValidateBasicInfoCases {
		for _, conditionCase := range errorValidateConditionCases {
			t.Run("With correct info and error condition"+basicInfoCase.testName+conditionCase.testName, func(t *testing.T) {

				testCase := advertisement.AddAdvertisementRequest{
					Title:      basicInfoCase.advertisement.Title,
					StartAt:    basicInfoCase.advertisement.StartAt,
					EndAt:      basicInfoCase.advertisement.EndAt,
					Conditions: conditionCase.advertisement.Conditions,
				}

				jsonData, err := json.Marshal(testCase)
				if err != nil {
					t.Errorf("Error to jsonify data with error " + err.Error())
					return
				}

				router := gin.New()
				router.Use(MiddlewareAddAdvertisement())
				router.POST("/test", func(c *gin.Context) {
					c.String(http.StatusOK, "Passed")
				})

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))

				router.ServeHTTP(w, req)

				assert.Equal(t, http.StatusBadRequest, w.Code)
			})
		}
	}

	for _, basicInfoCase := range errorValidateBasicInfoCases {
		for _, conditionCase := range rightValidateConditionCases {
			t.Run("With error info and correct condition"+basicInfoCase.testName+conditionCase.testName, func(t *testing.T) {

				testCase := advertisement.AddAdvertisementRequest{
					Title:      basicInfoCase.advertisement.Title,
					StartAt:    basicInfoCase.advertisement.StartAt,
					EndAt:      basicInfoCase.advertisement.EndAt,
					Conditions: conditionCase.advertisement.Conditions,
				}

				jsonData, err := json.Marshal(testCase)
				if err != nil {
					t.Errorf("Error to jsonify data with error " + err.Error())
					return
				}

				router := gin.New()
				router.Use(MiddlewareAddAdvertisement())
				router.POST("/test", func(c *gin.Context) {
					c.String(http.StatusOK, "Passed")
				})

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))

				router.ServeHTTP(w, req)

				assert.Equal(t, http.StatusBadRequest, w.Code)
			})
		}
	}

	for _, basicInfoCase := range errorValidateBasicInfoCases {
		for _, conditionCase := range errorValidateConditionCases {
			t.Run("With error info and error condition"+basicInfoCase.testName+conditionCase.testName, func(t *testing.T) {

				testCase := advertisement.AddAdvertisementRequest{
					Title:      basicInfoCase.advertisement.Title,
					StartAt:    basicInfoCase.advertisement.StartAt,
					EndAt:      basicInfoCase.advertisement.EndAt,
					Conditions: conditionCase.advertisement.Conditions,
				}

				jsonData, err := json.Marshal(testCase)
				if err != nil {
					t.Errorf("Error to jsonify data with error " + err.Error())
					return
				}

				router := gin.New()
				router.Use(MiddlewareAddAdvertisement())
				router.POST("/test", func(c *gin.Context) {
					c.String(http.StatusOK, "Passed")
				})

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))

				router.ServeHTTP(w, req)

				assert.Equal(t, http.StatusBadRequest, w.Code)
			})
		}
	}

}
