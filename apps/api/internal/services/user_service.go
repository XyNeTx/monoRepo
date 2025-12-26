package services

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fiber-api/internal/config"
	"fiber-api/internal/models"
	"fiber-api/internal/models/dto"
	"log"

	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

type userService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{DB: db}
}

func (us *userService) GetAllUsers() (userArr []dto.UserDTO, err error) {
	var users []models.User
	if err := us.DB.Find(&users).Select("email", "name", "surname", "age", "Address").Error; err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	userArr = make([]dto.UserDTO, 0, len(users))
	for _, user := range users {
		userArr = append(userArr, dto.UserDTO{
			Email:   user.Email,
			Name:    user.Name,
			Surname: user.Surname,
			Age:     user.Age,
			Address: user.Address,
		})

		//fmt.Printf("%+v\n", user)
	}
	//database.DB.Model(&models.User{}).Raw("SELECT * FROM users WHERE id = ?", 1).Scan(&userArr)
	return userArr, nil
}

func (us *userService) GetUserByEmail(email string) (dto.UserDTO, error) {
	var user dto.UserDTO
	result := us.DB.Where("email = ?", email).First(new(models.User)).Scan(&user)

	return user, result.Error
}

func (us *userService) GetUserById(id int) (dto.UserDTO, error) {
	var user dto.UserDTO
	result := us.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (us *userService) CreateUsers(userDTO dto.UserDTO) error {
	salt := make([]byte, 16)
	rand.Read(salt)

	pepper := config.GetEnv("HASH_SECRET")
	mac := hmac.New(sha256.New, []byte(pepper))
	if pepper == "" {
		log.Fatal("HASH_SECRET is not set")
	}
	mac.Write([]byte(userDTO.Password))
	pwdWithPepper := mac.Sum(nil)
	//fullPassword := string(pwdWithPepper)
	//fmt.Printf("FullPassword: %s\n", fullPassword)

	hash := argon2.IDKey(pwdWithPepper, salt, 1, 64*1024, 4, 32)
	hashB64 := base64.RawStdEncoding.EncodeToString(hash)
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)

	userObj := models.User{
		Email:        userDTO.Email,
		Name:         userDTO.Name,
		Surname:      userDTO.Surname,
		Age:          userDTO.Age,
		Address:      userDTO.Address,
		PasswordHash: hashB64,
		PasswordSalt: saltB64,
	}

	err := us.DB.Create(&userObj).Error
	return err

}

func (us *userService) EditUsers(userDTO dto.UserDTO) (dto.UserDTO, error) {
	var userDB models.User

	us.DB.Model(&models.User{}).Where("email = ?", userDTO.Email).First(&userDB)

	if userDB.ID == 0 {
		return userDTO, gorm.ErrRecordNotFound
	}

	userDB.Name = userDTO.Name
	userDB.Surname = userDTO.Surname
	userDB.Age = userDTO.Age
	userDB.Address = userDTO.Address

	result := us.DB.Save(&userDB)
	return userDTO, result.Error
}

func (us *userService) DeleteUser(email string) error {
	result := us.DB.Where("email = ?", email).Delete(&models.User{})
	//result = us.DB.Commit().Delete(&models.User{}, "email = ?", email)
	return result.Error
}

func (us *userService) VerifyPassword(email string, password string) (dto.UserDTO, error) {
	//userObj, err := GetUserByEmail(email)
	var userObj models.User
	us.DB.Where("email = ?", email).First(&userObj)
	if userObj.ID == 0 {
		return dto.UserDTO{}, errors.New("user not found")
	}
	//fmt.Printf("UserObj: %+v\n", userObj)

	pepper := config.GetEnv("HASH_SECRET")
	mac := hmac.New(sha256.New, []byte(pepper))
	if pepper == "" {
		log.Fatal("HASH_SECRET is not set")
	}
	mac.Write([]byte(password))
	pwdWithPepper := mac.Sum(nil)
	//fmt.Printf("FullPassword: %s\n", fullPassword)

	salt, _ := base64.RawStdEncoding.DecodeString(userObj.PasswordSalt)

	receiveHash := argon2.IDKey(pwdWithPepper, salt, 1, 64*1024, 4, 32)

	//fmt.Printf("ReceiveHash: %s\n", receiveHash)
	expectedHash, _ := base64.RawStdEncoding.DecodeString(userObj.PasswordHash)
	//fmt.Printf("ExpectedHash: %s\n", expectedHash)
	isValid := hmac.Equal(receiveHash, expectedHash)
	//fmt.Printf("IsValid: %v\n", isValid)

	if !isValid {
		return dto.UserDTO{}, errors.New("invalid password")
	}

	userDTO := dto.UserDTO{
		Email:   userObj.Email,
		Name:    userObj.Name,
		Surname: userObj.Surname,
		Age:     userObj.Age,
		Address: userObj.Address,
	}

	return userDTO, nil
}

func (us *userService) ChangePassword(email string, newPassword string) error {
	var userObj models.User
	us.DB.Where("email = ?", email).First(&userObj)

	if userObj == (models.User{}) {
		return gorm.ErrRecordNotFound
	}

	salt := make([]byte, 16)
	rand.Read(salt)
	pepper := config.GetEnv("HASH_SECRET")
	mac := hmac.New(sha256.New, []byte(pepper))
	mac.Write([]byte(newPassword))
	pwdWithPepper := mac.Sum(nil)

	hash := argon2.IDKey(pwdWithPepper, salt, 1, 64*1024, 4, 32)

	hashB64 := base64.RawStdEncoding.EncodeToString(hash)
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)

	userObj.PasswordHash = hashB64
	userObj.PasswordSalt = saltB64

	result := us.DB.Save(&userObj)
	return result.Error

}
