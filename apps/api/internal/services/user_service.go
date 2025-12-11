package services

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/base64"
	"fiber-api/internal/config"
	"fiber-api/internal/database"
	"fiber-api/internal/models"

	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

func GetAllUsers() (userArr []models.User, err error) {
	database.DB.Find(&userArr)
	//database.DB.Model(&models.User{}).Raw("SELECT * FROM users WHERE id = ?", 1).Scan(&userArr)
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
	salt := make([]byte, 16)
	rand.Read(salt)

	pepper := config.GetEnv("HASH_SECRET")
	fullPassword := userObj.PasswordHash + pepper

	hash := argon2.IDKey([]byte(fullPassword), salt, 1, 64*1024, 4, 32)

	saltB64 := base64.RawStdEncoding.EncodeToString(salt)
	hashB64 := base64.RawStdEncoding.EncodeToString(hash)

	userObj.PasswordHash = hashB64
	userObj.PasswordSalt = saltB64

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

func VerifyPassword(email string, password string) (bool, error) {
	userObj, err := GetUserByEmail(email)

	if err != nil {
		return false, err
	}

	pepper := config.GetEnv("HASH_SECRET")
	fullPassword := password + pepper

	salt, _ := base64.RawStdEncoding.DecodeString(userObj.PasswordSalt)

	receiveHash := argon2.IDKey([]byte(fullPassword), salt, 1, 64*1024, 4, 32)
	expectedHash, _ := base64.RawStdEncoding.DecodeString(userObj.PasswordHash)

	return hmac.Equal(receiveHash, expectedHash), nil
}
