/*
This is a tool to get the full url including parameter in query
*/

package utils

import "github.com/gin-gonic/gin"

// GetFullUrl is to get the full url and all query's parameters
func GetFullUrl(c *gin.Context) string {
	// return the combination of route and query parameter
	return c.Request.URL.RawPath + c.Request.URL.RawQuery
}
