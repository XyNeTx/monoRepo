package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"    // or the Docker service name if running in another container
	port     = "5432"         // default PostgreSQL port
	user     = "postgres"     // as defined in docker-compose.yml
	password = "@!Xn13799173" // as defined in docker-compose.yml
	dbname   = "postgres"     // as defined in docker-compose.yml
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect Database")
	}
	DB = db
}

// func SelectUser(db *gorm.DB, Email string) User {
// 	user := models.User{}

// 	db.Model(&models.User{}).Where("Email = ?", Email).Find(&user)

// 	return user
// }
