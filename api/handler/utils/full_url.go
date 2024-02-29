package utils

import "github.com/gin-gonic/gin"

func GetFullUrl(c *gin.Context) string {
	return c.Request.URL.RawPath + c.Request.URL.RawQuery
}
