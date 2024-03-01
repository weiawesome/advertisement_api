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
)

// InitRoutes is to initialize all routes and setting
func InitRoutes() *gin.Engine {
	// get the basic gin engine
	r := gin.Default()

	// set trust proxy to nil. if failed, log it and rerun nil.
	err := r.SetTrustedProxies(nil)
	if err != nil {
		utils.LogFatal(err.Error())
		return nil
	}

	// load the template files
	r.LoadHTMLGlob("templates/*")

	// register basic router
	basicRouter := r.Group("/api/" + utils.GetVersion())

	// register router for api documents
	apiDocsRouter := basicRouter.Group("/docs")
	// register router for advertisement affair
	advertisementRouter := basicRouter.Group("/ad")

	// new a sql repository
	sqlRepository := sql.NewRepository()
	// new a redis repository
	redisRepository := redis.NewRepository()

	// initialize api document router
	InitAPIDocsRoutes(apiDocsRouter)
	// initialize advertisement router
	InitAdvertisementRoutes(advertisementRouter, sqlRepository, redisRepository)

	// return the gin engine
	return r
}
