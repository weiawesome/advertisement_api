package utils

import (
	"advertisement_api/internal/repository/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB() (*gorm.DB, error) {
	user := EnvMySqlUser()
	password := EnvMySqlPassword()
	address := EnvMySqlAddress()
	dbName := EnvMySqlDb()
	location := GetSqlDsnLocation()

	dsn := user + ":" + password + "@tcp(" + address + ")/" + dbName + "?charset=utf8mb4&loc=" + location + "&parseTime=True"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func InitDB() error {
	var err error
	if db, err = connectDB(); err != nil {
		return err
	}
	if err := db.AutoMigrate(model.Advertisement{}, model.AgeCondition{}, model.CountryCondition{}, model.PlatformCondition{}, model.GenderCondition{}); err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
