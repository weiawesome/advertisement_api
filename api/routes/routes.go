package routes

import (
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}

	r.LoadHTMLGlob("templates/*")

	basicRouter := r.Group("/api/" + utils.GetVersion())

	apiDocsRouter := basicRouter.Group("/docs")
	advertisementRouter := basicRouter.Group("/ad")

	sqlRepository := sql.NewRepository()
	redisRepository := redis.NewRepository()

	InitAPIDocsRoutes(apiDocsRouter)
	InitAdvertisementRoutes(advertisementRouter, sqlRepository, redisRepository)

	return r
}
