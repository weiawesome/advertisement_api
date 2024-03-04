package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDB(t *testing.T) {
	t.Run("Right case get sql database", func(t *testing.T) {
		sqlDb := GetDB()
		assert.Equal(t, db, sqlDb)
	})
}

func TestCloseDB(t *testing.T) {
	t.Run("Right case when db is nil", func(t *testing.T) {
		err := CloseDB()
		assert.Nil(t, err)
	})
}
