package dto

type UserDTO struct {
	//gorm.Model
	Email    string `json:"email" gorm:"unique" example:"sitthiporn.po@gmail.com" required:"true"`
	Name     string `json:"name" example:"Sitthiporn" required:"true"`
	Surname  string `json:"surname" example:"Polmart" required:"true"`
	Age      int    `json:"age" example:22 required:"true"`
	Address  string `json:"address" example:"123/45 Moo 6, T. Bangna, A. Bangna, Bangkok 10260"`
	Password string `json:"password"`
}
