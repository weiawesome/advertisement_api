package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConnectDB(t *testing.T) {
	t.Run("Right case", func(t *testing.T) {
		host := os.Getenv("MYSQL_HOST")
		err := os.Setenv("MYSQL_HOST", "Not exist host")
		if err != nil {
			t.Errorf("error to get environment " + err.Error())
			return
		}
		_, err = connectDB()
		assert.NotNil(t, err)
		err = os.Setenv("MYSQL_HOST", host)
		if err != nil {
			t.Errorf("error to set environment " + err.Error())
			return
		}
	})
}

func TestInitDB(t *testing.T) {
	t.Run("Right case", func(t *testing.T) {
		host := os.Getenv("MYSQL_HOST")
		err := os.Setenv("MYSQL_HOST", "Not exist host")
		if err != nil {
			t.Errorf("error to get environment " + err.Error())
			return
		}
		err = InitDB()
		assert.NotNil(t, err)
		err = os.Setenv("MYSQL_HOST", host)
		if err != nil {
			t.Errorf("error to set environment " + err.Error())
			return
		}
	})
}

func TestGetDB(t *testing.T) {
	t.Run("Right case get sql database", func(t *testing.T) {
		sqlDb := GetDB()
		assert.Equal(t, db, sqlDb)
	})
}

func TestCloseDB(t *testing.T) {
	t.Run("Right case when db is nil", func(t *testing.T) {
		db = nil
		err := CloseDB()
		assert.Nil(t, err)
	})
	t.Run("Right case when db is not nil", func(t *testing.T) {
		_ = InitDB()
		err := CloseDB()
		assert.Nil(t, err)
	})
}
