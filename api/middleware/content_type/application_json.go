/*
The middleware of gin to check content-type is application/json or not
*/

package content_type

import (
	"advertisement_api/api/response/failure"
	"github.com/gin-gonic/gin"
	"net/http"
)

// MiddlewareApplicationJson to check the content-type
func MiddlewareApplicationJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the content-type
		contentType := c.ContentType()

		// when content-type is not application/json then return error
		if contentType != "application/json" {
			e := failure.ClientError{Reason: "Content-Type must be application/json not " + contentType}
			c.JSON(http.StatusUnsupportedMediaType, e)
			c.Abort()
			return
		}

		// go to the next handler
		c.Next()
	}
}
