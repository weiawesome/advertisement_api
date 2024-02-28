package content_type

import (
	"advertisement_api/api/response/failure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddlewareApplicationJson() gin.HandlerFunc {
	return func(c *gin.Context) {

		contentType := c.ContentType()

		if contentType != "application/json" {
			e := failure.ClientError{Reason: "Content-Type must be application/json not " + contentType}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		c.Next()
	}
}
