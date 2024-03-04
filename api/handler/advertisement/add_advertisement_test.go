package advertisement

import (
	advertisement2 "advertisement_api/api/request/advertisement"
	"advertisement_api/api/response/failure"
	"advertisement_api/internal/service/advertisement"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	normalCase        = advertisement2.AddAdvertisementRequest{Title: func(s string) *string { return &s }(advertisement.NormalTitleMock)}
	dayLimitCase      = advertisement2.AddAdvertisementRequest{Title: func(s string) *string { return &s }(advertisement.DayLimitTitleMock)}
	durationLimitCase = advertisement2.AddAdvertisementRequest{Title: func(s string) *string { return &s }(advertisement.DurationLimitTitleMock)}
	unknownErrorCase  = advertisement2.AddAdvertisementRequest{Title: func(s string) *string { return &s }(advertisement.UnknownErrorTitleMock)}
)

func TestAddAdvertisementHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Case right", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.AddAdvertisementServiceMock)
		handler := HandlerAddAdvertisement{Service: mockService}

		router.POST("/test",
			func(c *gin.Context) { c.Set("data", normalCase) },
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Case error with day limitation exceeded", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.AddAdvertisementServiceMock)
		handler := HandlerAddAdvertisement{Service: mockService}

		router.POST("/test",
			func(c *gin.Context) { c.Set("data", dayLimitCase) },
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusTooManyRequests, w.Code)
		var e failure.DayLimitError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Day limit error, "+*dayLimitCase.Title, e.Reason)
	})
	t.Run("Case error with duration limitation exceeded", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.AddAdvertisementServiceMock)
		handler := HandlerAddAdvertisement{Service: mockService}

		router.POST("/test",
			func(c *gin.Context) { c.Set("data", durationLimitCase) },
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
		var e failure.DurationLimitError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Duration limit error, "+*durationLimitCase.Title, e.Reason)
	})
	t.Run("Case error with unknown reason", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.AddAdvertisementServiceMock)
		handler := HandlerAddAdvertisement{Service: mockService}

		router.POST("/test",
			func(c *gin.Context) { c.Set("data", unknownErrorCase) },
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Server error, "+*unknownErrorCase.Title, e.Reason)
	})
	t.Run("Case error with without set in context", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.AddAdvertisementServiceMock)
		handler := HandlerAddAdvertisement{Service: mockService}

		router.POST("/test",
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Data not found in context", e.Reason)
	})
}
