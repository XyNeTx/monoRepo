package models

import (
	"gorm.io/gorm"
)

type User struct {
	//gorm.Model
	ID           int `gorm:"primaryKey" example:"8"`
	created_at   gorm.DeletedAt 
	updated_at   gorm.DeletedAt
	deleted_at   gorm.DeletedAt
	Email        string `gorm:"unique" example:"sitthiporn.po@gmail.com"`
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
