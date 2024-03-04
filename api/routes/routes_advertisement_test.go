/*
There is initialization for the advertisement affairs.
Including add advertisement and get advertisements.
Lots of middleware to check content and handler with service to process the request.
*/

package routes

import (
	"advertisement_api/api/request/advertisement"
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"advertisement_api/utils"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

var (
	emptyCase = ""

	ageRightCases = []string{
		emptyCase, strconv.Itoa(utils.GetMinAge()), "24", strconv.Itoa(utils.GetMaxAge()),
	}
	countryRightCases = []string{
		emptyCase, "TW", "JP",
	}
	genderRightCases = []string{
		emptyCase, "M", "F",
	}
	platformRightCases = []string{
		emptyCase, "ios", "android", "web",
	}
	offsetRightCases = []string{
		emptyCase, strconv.Itoa(utils.GetMinOffset()), "10",
	}
	limitRightCases = []string{
		emptyCase, strconv.Itoa(utils.GetMinLimit()), "10",
	}

	advertisementRightCases = []advertisement.AddAdvertisementRequest{
		{
			Title:   func(s string) *string { return &s }(sql.NormalCase),
			StartAt: func(t time.Time) *time.Time { return &t }(time.Now().UTC()),
			EndAt:   func(t time.Time) *time.Time { return &t }(time.Now().UTC().Add(time.Hour * 10)),
		},
	}

	contentTypeRightCase = "application/json"
)

func TestInitAdvertisementRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	utils.InitMaps()
	utils.InitSingleFlight()

	for _, age := range ageRightCases {
		for _, country := range countryRightCases {
			for _, gender := range genderRightCases {
				for _, platform := range platformRightCases {
					for _, offset := range offsetRightCases {
						for _, limit := range limitRightCases {
							t.Run("Right case", func(t *testing.T) {
								r := gin.New()
								group := r.Group("/ad")

								sqlRepo := new(sql.RepositoryMock)
								redisRepo := new(redis.RepositoryMock)

								InitAdvertisementRoutes(group, sqlRepo, redisRepo)

								w := httptest.NewRecorder()

								queryRoute := "/ad?"

								if age != emptyCase {
									queryRoute += "age=" + age + "&"
								}
								if country != emptyCase {
									queryRoute += "country=" + country + "&"
								}
								if gender != emptyCase {
									queryRoute += "gender=" + gender + "&"
								}
								if platform != emptyCase {
									queryRoute += "platform=" + platform + "&"
								}
								if offset != emptyCase {
									queryRoute += "offset=" + offset + "&"
								}
								if limit != emptyCase {
									queryRoute += "limit=" + limit + "&"
								}

								queryRoute = queryRoute[:len(queryRoute)-1]

								req, _ := http.NewRequest("GET", queryRoute, nil)
								req.Header.Set("Content-Type", contentTypeRightCase)
								r.ServeHTTP(w, req)
								assert.Equal(t, 200, w.Code, "GET request should respond with status 200")
							})
						}
					}
				}
			}
		}
	}

	for _, rightCase := range advertisementRightCases {
		t.Run("Right case", func(t *testing.T) {
			r := gin.New()
			group := r.Group("/ad")

			sqlRepo := new(sql.RepositoryMock)
			redisRepo := new(redis.RepositoryMock)

			InitAdvertisementRoutes(group, sqlRepo, redisRepo)

			w := httptest.NewRecorder()
			jsonData, err := json.Marshal(rightCase)
			if err != nil {
				t.Errorf("Error to jsonify data with error " + err.Error())
				return
			}

			req, _ := http.NewRequest("POST", "/ad", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", contentTypeRightCase)
			r.ServeHTTP(w, req)
			assert.Equal(t, 200, w.Code, "GET request should respond with status 200")
		})
	}

}
