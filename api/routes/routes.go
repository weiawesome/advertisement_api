package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}

	r.LoadHTMLGlob("templates/*")

	basicRouter := r.Group("/api/v1")

	apiDocsRouter := basicRouter.Group("/docs")

	InitAPIDocsRoutes(apiDocsRouter)

	return r
}
