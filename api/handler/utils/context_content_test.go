package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type objectCase struct{}

var (
	testIntCase    = 10
	testStrCase    = "Test"
	testBoolCase   = true
	testFloatCase  = 1.5
	testSliceCase  = []string{"test"}
	testChanCase   = make(chan int, 5)
	testMapCase    = map[int]bool{1: true}
	testObjectCase = objectCase{}

	notExistCase = ""
)

func TestGetFromContext(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Case right with integer", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			c.Set("key", testIntCase)
			value, ok := GetFromContext[int](c, "key")
			assert.True(t, ok)
			assert.Equal(t, testIntCase, value)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})
	t.Run("Case right with string", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			c.Set("key", testStrCase)
			value, ok := GetFromContext[string](c, "key")
			assert.True(t, ok)
			assert.Equal(t, testStrCase, value)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})
	t.Run("Case right with boolean", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			c.Set("key", testBoolCase)
			value, ok := GetFromContext[bool](c, "key")
			assert.True(t, ok)
			assert.Equal(t, testBoolCase, value)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})
	t.Run("Case right with float", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			c.Set("key", testFloatCase)
			value, ok := GetFromContext[float64](c, "key")
			assert.True(t, ok)
			assert.Equal(t, testFloatCase, value)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})
	t.Run("Case right with slice", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			c.Set("key", testSliceCase)
			value, ok := GetFromContext[[]string](c, "key")
			assert.True(t, ok)
			assert.Equal(t, testSliceCase, value)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})
	t.Run("Case right with chan", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			c.Set("key", testChanCase)
			value, ok := GetFromContext[chan int](c, "key")
			assert.True(t, ok)
			assert.Equal(t, testChanCase, value)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})
	t.Run("Case right with map", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			c.Set("key", testMapCase)
			value, ok := GetFromContext[map[int]bool](c, "key")
			assert.True(t, ok)
			assert.Equal(t, testMapCase, value)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})
	t.Run("Case right with structure", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			c.Set("key", testObjectCase)
			value, ok := GetFromContext[objectCase](c, "key")
			assert.True(t, ok)
			assert.Equal(t, testObjectCase, value)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})

	t.Run("Case error with not exist", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", func(c *gin.Context) {
			_, ok := GetFromContext[string](c, notExistCase)
			assert.False(t, ok)
		})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		r.ServeHTTP(w, req)
	})
}
