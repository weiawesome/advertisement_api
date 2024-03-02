package content_type

import (
	"advertisement_api/api/response/failure"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	rightCase  = "application/json"
	errorCases = []string{
		"text/plain", "text/html", "text/xml",
		"application/xml", "application/javascript", "application/octet-stream", "application/pdf",
		"multipart/form-data",
		"image/jpeg", "image/png", "image/gif",
	}
)

func TestMiddlewareApplicationJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("With correct Content-Type", func(t *testing.T) {
		router := gin.New()
		router.Use(MiddlewareApplicationJson())
		router.POST("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "Passed")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", nil)
		req.Header.Set("Content-Type", rightCase)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Passed", w.Body.String())
	})

	for _, errorCase := range errorCases {
		t.Run("With incorrect Content-Type "+errorCase, func(t *testing.T) {
			router := gin.New()
			router.Use(MiddlewareApplicationJson())
			router.POST("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Passed")
			})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/test", nil)
			req.Header.Set("Content-Type", errorCase)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)

			var e failure.ClientError
			err := json.Unmarshal(w.Body.Bytes(), &e)
			if err != nil {
				t.Fatal("Failed to unmarshal response body:", err)
			}
			assert.Equal(t, "Content-Type must be application/json not "+errorCase, e.Reason)
		})
	}
}
