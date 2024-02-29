package utils

import "github.com/gin-gonic/gin"

func GetFromContext[T any](c *gin.Context, key string) (T, bool) {
	val, ok := c.Get(key)
	if !ok {
		return *new(T), false
	}
	result, resultOk := val.(T)
	return result, resultOk
}
