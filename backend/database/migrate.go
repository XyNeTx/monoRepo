package database

import "fiber-api/models"

func Migrate() {
	DB.AutoMigrate(
		&models.User{},
	)
}
