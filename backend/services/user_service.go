package services

import (
	"fiber-api/database"
	"fiber-api/models"

	"gorm.io/gorm"
)

func GetAllUsers() (userArr []models.User, err error) {
	//database.DB.Find(&userArr)
	database.DB.Model(&models.User{}).Raw("SELECT * FROM users WHERE id = ?", 1).Scan(&userArr)
	return userArr, err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}

func CreateUsers(userObj models.User) (models.User, error) {
	// var userList []models.User

	// var createUser models.User
	// createUser.Age = 29
	// createUser.Email = "exzdue3@gmail.com"
	// createUser.Name = "Sitthiporn"
	// createUser.Surname = "Polmart"

	// createUser2 := models.User{
	// 	Age:     29,
	// 	Email:   "exzdue2@gmail.com",
	// 	Name:    "Sitthiporn",
	// 	Surname: "Polmart",
	// }

	// userList = append(userList, createUser, createUser2)
	// err = database.DB.Create(&userList).Error
	result := database.DB.Create(&userObj)
	return userObj, result.Error
}

func EditUsers(userObj models.User) (models.User, error) {
	var userDB models.User
	database.DB.Model(&models.User{}).Where("email = ?", userObj.Email).First(&userDB)
	if userDB.ID == 0 {
		return userDB, gorm.ErrRecordNotFound
	}
	userDB.Name = userObj.Name
	userDB.Surname = userObj.Surname
	userDB.Age = userObj.Age

	result := database.DB.Save(&userDB)
	return userDB, result.Error
}

func DeleteUser(email string) error {
	result := database.DB.Where("email = ?", email).Delete(&models.User{})
	return result.Error
}
