/*
There is initialization of the routes.
It will set some setting at first include trusted proxy, template files, groups routes etc.
Then new redis and sql repository. Finally, initialize all routes and return it.
*/

package routes

import (
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"runtime"
)

// InitRoutes is to initialize all routes and setting
func InitRoutes(sqlRepository sql.Repository, redisRepository redis.Repository) *gin.Engine {
	// get the basic gin engine
	r := gin.Default()

	// get the current file's path
	_, currentFile, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(currentFile)

	// load the template files
	r.LoadHTMLGlob(filepath.Join(basePath, "../../templates/*"))

	// register basic router
	basicRouter := r.Group("/api/" + utils.GetVersion())

	// register router for api documents
	apiDocsRouter := basicRouter.Group("/docs")
	// register router for advertisement affair
	advertisementRouter := basicRouter.Group("/ad")

	// initialize api document router
	InitAPIDocsRoutes(apiDocsRouter)
	// initialize advertisement router
	InitAdvertisementRoutes(advertisementRouter, sqlRepository, redisRepository)

	// return the gin engine
	return r
}
