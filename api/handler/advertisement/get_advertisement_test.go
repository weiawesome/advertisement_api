/*
The handler with service to handle about get advertisements.
At first, it will get value form context then throw the value to the service.
Finally, according to the value rerun from service then return different response to the client.
*/

package advertisement

import (
	"advertisement_api/api/response/failure"
	"advertisement_api/internal/service/advertisement"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetAdvertisementCase struct {
	Key      string
	Age      int
	Gender   string
	Country  string
	Platform string
	Offset   int
	Limit    int
}

var (
	normalCountryCase       = GetAdvertisementCase{Country: advertisement.NormalCountryMock}
	unknownErrorCountryCase = GetAdvertisementCase{Country: advertisement.UnknownErrorCountryMock}
)

// Handle to handle the request about getting advertisement
func TestHandlerGetAdvertisementHandle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Case right", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.GetAdvertisementServiceMock)
		handler := HandlerGetAdvertisement{Service: mockService}

		router.GET("/test",
			func(c *gin.Context) {
				c.Set("age", normalCountryCase.Age)
				c.Set("country", normalCountryCase.Country)
				c.Set("gender", normalCountryCase.Gender)
				c.Set("platform", normalCountryCase.Platform)
				c.Set("offset", normalCountryCase.Offset)
				c.Set("limit", normalCountryCase.Limit)
			},
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("Case error with unknown reason", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.GetAdvertisementServiceMock)
		handler := HandlerGetAdvertisement{Service: mockService}

		router.GET("/test",
			func(c *gin.Context) {
				c.Set("age", unknownErrorCountryCase.Age)
				c.Set("country", unknownErrorCountryCase.Country)
				c.Set("gender", unknownErrorCountryCase.Gender)
				c.Set("platform", unknownErrorCountryCase.Platform)
				c.Set("offset", unknownErrorCountryCase.Offset)
				c.Set("limit", unknownErrorCountryCase.Limit)
			},
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Server error, "+*unknownErrorCase.Title, e.Reason)
	})
	t.Run("Case error with without set in context(age)", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.GetAdvertisementServiceMock)
		handler := HandlerGetAdvertisement{Service: mockService}

		router.GET("/test",
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Age not found in context", e.Reason)
	})
	t.Run("Case error with without set in context(country)", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.GetAdvertisementServiceMock)
		handler := HandlerGetAdvertisement{Service: mockService}

		router.GET("/test",
			func(c *gin.Context) {
				c.Set("age", normalCountryCase.Age)
			},
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Country not found in context", e.Reason)
	})
	t.Run("Case error with without set in context(gender)", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.GetAdvertisementServiceMock)
		handler := HandlerGetAdvertisement{Service: mockService}

		router.GET("/test",
			func(c *gin.Context) {
				c.Set("age", normalCountryCase.Age)
				c.Set("country", normalCountryCase.Country)
			},
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Gender not found in context", e.Reason)
	})
	t.Run("Case error with without set in context(platform)", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.GetAdvertisementServiceMock)
		handler := HandlerGetAdvertisement{Service: mockService}

		router.GET("/test",
			func(c *gin.Context) {
				c.Set("age", normalCountryCase.Age)
				c.Set("country", normalCountryCase.Country)
				c.Set("gender", normalCountryCase.Gender)
			},
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Platform not found in context", e.Reason)
	})
	t.Run("Case error with without set in context(offset)", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.GetAdvertisementServiceMock)
		handler := HandlerGetAdvertisement{Service: mockService}

		router.GET("/test",
			func(c *gin.Context) {
				c.Set("age", normalCountryCase.Age)
				c.Set("country", normalCountryCase.Country)
				c.Set("gender", normalCountryCase.Gender)
				c.Set("platform", normalCountryCase.Platform)
			},
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Offset not found in context", e.Reason)
	})
	t.Run("Case error with without set in context(limit)", func(t *testing.T) {
		router := gin.New()
		mockService := new(advertisement.GetAdvertisementServiceMock)
		handler := HandlerGetAdvertisement{Service: mockService}

		router.GET("/test",
			func(c *gin.Context) {
				c.Set("age", normalCountryCase.Age)
				c.Set("country", normalCountryCase.Country)
				c.Set("gender", normalCountryCase.Gender)
				c.Set("platform", normalCountryCase.Platform)
				c.Set("offset", normalCountryCase.Offset)
			},
			handler.Handle,
		)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var e failure.ServerError
		err := json.Unmarshal(w.Body.Bytes(), &e)
		if err != nil {
			t.Fatal("Failed to unmarshal response body:", err)
		}
		assert.Equal(t, "Limit not found in context", e.Reason)
	})
}
