/*
There is a gorm db instance connect with mysql database.
If program start, it will initialize the instance and try to connect mysql database.
Furthermore, get function and close function is to get mysql client and close mysql connection.
*/

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetDB is to get the db instance
func TestGetDB(t *testing.T) {
	t.Run("Right case get sql database", func(t *testing.T) {
		sqlDb := GetDB()
		assert.Equal(t, db, sqlDb)
	})
}

// CloseDB is to close the mysql database connection
func TestCloseDB(t *testing.T) {
	t.Run("Right case when db is nil", func(t *testing.T) {
		err := CloseDB()
		assert.Nil(t, err)
	})
}
