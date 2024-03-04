package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestInitAPIDocsRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Case with OpenAPI(SWAGGER) html file", func(t *testing.T) {
		router := gin.New()
		router.LoadHTMLGlob("../../templates/*")
		group := router.Group("/docs")

		InitAPIDocsRoutes(group)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/docs", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
	t.Run("Case with static file", func(t *testing.T) {
		router := gin.New()

		group := router.Group("/docs")

		InitAPIDocsRoutes(group)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/docs/api_specification.yaml", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
	})
}
