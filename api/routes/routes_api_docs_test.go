/*
There is initialization for the api document affairs.
Set the static file (api document file) and manifest it on the html.
*/

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

// InitAPIDocsRoutes is to initialize the api document route
func TestInitAPIDocsRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Case with OpenAPI(SWAGGER) html file", func(t *testing.T) {
		router := gin.New()
		// 設置HTML渲染器
		router.LoadHTMLGlob("../../templates/*")
		group := router.Group("/docs")

		InitAPIDocsRoutes(group)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/docs", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code) // 驗證狀態碼，假設index.html存在且可訪問
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
