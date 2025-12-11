package database

import (
	"fiber-api/internal/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		config.GetEnv("DB_HOST"), config.GetEnv("DB_PORT"), config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASS"), config.GetEnv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect Database")
	}

	DB = db
}
