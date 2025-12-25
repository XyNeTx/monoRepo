package controllers

import (
	"fiber-api/internal/models"
	"fiber-api/internal/models/dto"
	"fiber-api/internal/services/interfaces"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type UserController struct {
	service interfaces.UserService
}

func NewUserController(service interfaces.UserService) *UserController {
	return &UserController{service: service}
}

// @Summary Welcome to Fiber API
// @Tags Root
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router / [get]
func InitialUserController(c fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "Welcome to Fiber API"})
}

// @Summary Get All Users
// @Tags Users
// @Accept json
// @Produce json“
// @Success 200 {array} models.User
// @Failure 500 "Internal Server Error"
// @Router /api/users [get]
func (uc *UserController) GetUsers(c fiber.Ctx) error {
	userArr, err := uc.service.GetAllUsers()

	if userArr == nil {
		return c.Status(404).JSON(fiber.Map{"message": "No users found"})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to retrieve users"})
	}

	return c.Status(200).JSON(fiber.Map{"data": userArr})
}

// @Summary Get User by Email
// @Tags Users
// @Accept json
// @Produce json
// @Param	email Sitthiporn.po@gmail.coml	path	string	true  "The email of the resource Sitthiporn.po@gmail.com"
// @Description Get a user by their email address
// @Success 200 {object} models.User
// @Failure 401 "Unauthorized"
// @Router /api/users/email/{email} [get]
func (uc *UserController) GetUserByEmail(c fiber.Ctx) error {
	email := c.Params("email")
	user, err := uc.service.GetUserByEmail(email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to retrieve user", "error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"data": user})
}

// @Summary Get All Users
// @Tags Users
// @Accept json
// @Produce json
// @Param	id	path	string	true "The ID of the user"
// @Success 200 {array} models.User
// @Failure 500 "Internal Server Error"
// @Router /api/users/id/{id} [get]
func (uc *UserController) GetUserById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid ID"})
	}
	user, err := uc.service.GetUserById(id)
	if user == (dto.UserDTO{}) {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to retrieve user", "error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"data": user})
}

func (uc *UserController) CreateUsers(c fiber.Ctx) error {
	user := new(dto.UserDTO)

	if err := c.Bind().JSON(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	err := uc.service.CreateUsers(*user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to create users"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Success"})
}

func (uc *UserController) EditUsers(c fiber.Ctx) error {
	user := new(dto.UserDTO)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	editedUser, err := uc.service.EditUsers(*user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to edit user"})
	}

	return c.Status(200).JSON(fiber.Map{"data": editedUser, "message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c fiber.Ctx) error {
	email := c.Params("email")
	err := uc.service.DeleteUser(email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to delete user"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "User deleted successfully"})
}

func (uc *UserController) LoginUser(c fiber.Ctx) error {
	loginReq := new(models.User)

	if err := c.Bind().Body(loginReq); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}
	isValid, err := uc.service.VerifyPassword(loginReq.Email, loginReq.PasswordHash)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Login failed"})
	}

	if !isValid {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid Email or Password"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Login successful"})
}

func (uc *UserController) ChangePassword(c fiber.Ctx) error {
	changePasswordReq := new(dto.UserDTO)

	if err := c.Bind().Body(changePasswordReq); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	err := uc.service.ChangePassword(changePasswordReq.Email, changePasswordReq.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Change password failed"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Change password successful"})
}
