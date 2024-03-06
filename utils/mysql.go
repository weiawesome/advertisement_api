/*
There is a gorm db instance connect with mysql database.
If program start, it will initialize the instance and try to connect mysql database.
Furthermore, get function and close function is to get mysql client and close mysql connection.
*/

package utils

import (
	"advertisement_api/internal/repository/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm db instance
var db *gorm.DB

// this is try to connect mysql database
func connectDB() (*gorm.DB, error) {
	// get the user, password, address, db name, location from environment
	user := EnvMySqlUser()
	password := EnvMySqlPassword()
	address := EnvMySqlAddress()
	dbName := EnvMySqlDb()
	location := GetSqlDsnLocation()

	// build the dsn and try to connect the database
	dsn := user + ":" + password + "@tcp(" + address + ")/" + dbName + "?charset=utf8mb4&loc=" + location + "&parseTime=True"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

// InitDB is to initialize the db connection
func InitDB() error {
	var err error

	// try to connect sql database
	if db, err = connectDB(); err != nil {
		return err
	}

	// try to migrate the model to the database
	// models including advertisement, age condition, country condition, platform condition and gender condition
	return db.AutoMigrate(model.Advertisement{}, model.AgeCondition{}, model.CountryCondition{}, model.PlatformCondition{}, model.GenderCondition{})
}

// GetDB is to get the db instance
func GetDB() *gorm.DB {
	return db
}

// CloseDB is to close the mysql database connection
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
