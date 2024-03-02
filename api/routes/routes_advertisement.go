/*
There is initialization for the advertisement affairs.
Including add advertisement and get advertisements.
Lots of middleware to check content and handler with service to process the request.
*/

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

// InitAdvertisementRoutes is to initialize the advertisement routes
func InitAdvertisementRoutes(r *gin.RouterGroup, sqlRepository *sql.Repository, redisRepository *redis.Repository) {
	// get method is to get the advertisements
	r.GET("", // additional route
		content_type.MiddlewareApplicationJson(),                               // check content-type
		getAdvertisement.MiddlewareAge(), getAdvertisement.MiddlewareCountry(), // check age and country
		getAdvertisement.MiddlewareGender(), getAdvertisement.MiddlewarePlatform(), // check gender and platform
		getAdvertisement.MiddlewarePaginationOffset(), getAdvertisement.MiddlewarePaginationLimit(), // check limit and offset for pagination
		(&advertisement.HandlerGetAdvertisement{
			Service: advertisementService.NewGetAdvertisementService(sqlRepository, redisRepository),
		}).Handle) // handler with service to process

	// post method is to add the advertisement by query content
	r.POST("", // additional route
		content_type.MiddlewareApplicationJson(),      // check content-type
		addAdvertisement.MiddlewareAddAdvertisement(), // check the request-body content
		(&advertisement.HandlerAddAdvertisement{
			Service: advertisementService.NewAddAdvertisementService(sqlRepository),
		}).Handle) // handler with service to process
}
