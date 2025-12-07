package services

import (
	"gin-api/database"
	"gin-api/models"
)

func GetAllUsers() (user []models.User, err error) {
	database.DB.Find(&user)
	return user, err
}

func CreateUsers() (err error) {
	var userList []models.User

	var createUser models.User
	createUser.Age = 29
	createUser.Email = "exzdue@gmail.com"
	createUser.Name = "Sitthiporn"
	createUser.Surname = "Polmart"

	createUser2 := models.User{
		Age:     29,
		Email:   "Nipaporn_Rupsom@gmail.com",
		Name:    "Nipaporn",
		Surname: "Rupsom",
	}

	userList = append(userList, createUser, createUser2)
	database.DB.Create(&userList)

	return err
}
