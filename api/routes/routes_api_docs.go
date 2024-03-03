/*
There is initialization for the api document affairs.
Set the static file (api document file) and manifest it on the html.
*/

package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"runtime"
)

// InitAPIDocsRoutes is to initialize the api document route
func InitAPIDocsRoutes(r *gin.RouterGroup) {
	// get the current file's path
	_, currentFile, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(currentFile)

	// set the static file in route
	r.StaticFile("/api_specification.yaml", filepath.Join(basePath, "../../assets/docs/api_specification.yaml"))

	// set the route to manifest the OpenAPI(SWAGGER)
	r.GET("", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}
