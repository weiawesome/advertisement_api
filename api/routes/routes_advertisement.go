package routes

import (
	"advertisement_api/api/handler/advertisement"
	"advertisement_api/api/middleware/content_type"
	addAdvertisement "advertisement_api/api/middleware/request_body/advertisement"
	getAdvertisement "advertisement_api/api/middleware/url_parameter/advertisement"
	"advertisement_api/internal/repository/redis"
	"advertisement_api/internal/repository/sql"
	advertisementService "advertisement_api/internal/service/advertisement"
	"github.com/gin-gonic/gin"
)

func InitAdvertisementRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	r.GET("", content_type.MiddlewareApplicationJson(),
		getAdvertisement.MiddlewareAge(), getAdvertisement.MiddlewareCountry(),
		getAdvertisement.MiddlewareGender(), getAdvertisement.MiddlewarePlatform(),
		getAdvertisement.MiddlewarePagination(),
		(&advertisement.HandlerGetAdvertisement{Service: advertisementService.ServiceGetAdvertisement{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}}).Handle)
	r.POST("", content_type.MiddlewareApplicationJson(),
		addAdvertisement.MiddlewareAddAdvertisement(),
		(&advertisement.HandlerAddAdvertisement{Service: advertisementService.ServiceAddAdvertisement{SqlRepository: *sqlRepository, RedisRepository: *redisRepository}}).Handle)
}
