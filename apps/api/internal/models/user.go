package models

import (
	"time"
)

type User struct {
	//gorm.Model
	ID           int       `gorm:"primaryKey" example:"8"`
	CreatedAt    time.Time // `json:"created_at"`
	UpdatedAt    time.Time // `json:"updated_at"`
	DeletedAt    time.Time // `json:"deleted_at"`
	Email        string    `json:"email" gorm:"unique" example:"sitthiporn.po@gmail.com" required:"true"`
	Name         string    `json:"name" example:"Sitthiporn" required:"true"`
	Surname      string    `json:"surname" example:"Polmart" required:"true"`
	Age          int       `json:"age" example:"22" required:"true"`
	Token        string    `json:"token"`
	Address      string    `json:"address" example:"123/45 Moo 6, T. Bangna, A. Bangna, Bangkok 10260"`
	PasswordHash string    `json:"password"`
	PasswordSalt string    `json:"-"`
}

// func (User) TableName() string {
// 	return "user"
// }
