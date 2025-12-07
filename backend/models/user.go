package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email   string `gorm:"unique"`
	Name    string
	Surname string
	Age     int
}

func (User) TableName() string {
	return "user"
}
