package routes

import (
	"advertisement_api/api/middleware/content_type"
	addAdvertisement "advertisement_api/api/middleware/request_body/advertisement"
	getAdvertisement "advertisement_api/api/middleware/url_parameter/advertisement"
	"github.com/gin-gonic/gin"
)

func InitAdvertisementRoutes(r *gin.RouterGroup) {
	r.GET("", content_type.MiddlewareApplicationJson(),
		getAdvertisement.MiddlewareAge(), getAdvertisement.MiddlewareCountry(),
		getAdvertisement.MiddlewareGender(), getAdvertisement.MiddlewarePlatform(),
		getAdvertisement.MiddlewarePagination())
	r.POST("", content_type.MiddlewareApplicationJson(),
		addAdvertisement.MiddlewareAddAdvertisement())
}
