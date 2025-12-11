package database

import "fiber-api/internal/models"

func Migrate() {
	DB.AutoMigrate(
		&models.User{},
	)
}
