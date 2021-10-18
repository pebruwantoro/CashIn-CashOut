package config

import (
	"cash-app/models"
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
	var err error
	config := map[string]string{
		"DB_USERNAME": os.Getenv("DB_USERNAME"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_NAME":     os.Getenv("DB_NAME"),
	}
	ConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"],
	)
	DB, err = gorm.Open(mysql.Open(ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitDBTest() (*gorm.DB, error) {
	var err error
	config := map[string]string{
		"DB_USERNAME": os.Getenv("DB_USERNAME"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_NAME":     os.Getenv("TESTING_DB_NAME"),
	}
	ConnectionStringTest := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"],
	)
	DB, err = gorm.Open(mysql.Open(ConnectionStringTest), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	return DB, err
}

func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
