/*
This is a tool to take the value from context and assert its type
*/

package utils

import "github.com/gin-gonic/gin"

// GetFromContext is a tool to get the context and assert its type
func GetFromContext[T any](c *gin.Context, key string) (T, bool) {
	// get the value from context
	val, ok := c.Get(key)

	// not get from context then return T and false
	if !ok {
		return *new(T), false
	}

	// assert the value to specified type T and return value
	result, resultOk := val.(T)
	return result, resultOk
}
