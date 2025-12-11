package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique"`
	Name         string
	Surname      string
	Age          int
	Token        string
	Address      string
	PasswordHash string
	PasswordSalt string
}

// func (User) TableName() string {
// 	return "user"
// }
