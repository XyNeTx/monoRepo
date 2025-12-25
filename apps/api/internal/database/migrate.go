package database

import (
	"fiber-api/internal/models"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(
		&models.User{},
	)
}
