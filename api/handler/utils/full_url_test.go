/*
This is a tool to get the full url including parameter in query
*/

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testCases = []string{
		"/test",
		"/test?age=24",
		"/test?gender=M",
		"/test?platform=ios",
		"/test?country=TW",
		"/test?age=24&gender=M",
		"/test?age=24&gender=M&platform=ios",
		"/test?age=24&gender=M&platform=ios&country=TW",
	}
)

// GetFullUrl is to get the full url and all query's parameters
func TestGetFullUrl(t *testing.T) {

	for _, testCase := range testCases {
		t.Run("Case with "+testCase, func(t *testing.T) {
			r := gin.New()

			r.GET("/test", func(c *gin.Context) {
				fullUrl := GetFullUrl(c)
				c.String(http.StatusOK, fullUrl)
			})

			req := httptest.NewRequest("GET", testCase, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, testCase, w.Body.String())
		})
	}
}
