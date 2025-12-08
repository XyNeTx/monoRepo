package controllers

import (
	"fiber-api/models"
	"fiber-api/services"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	userArr, err := services.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to retrieve users"})
	}
	return c.Status(200).JSON(fiber.Map{"data": userArr})
}

func GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := services.GetUserByEmail(email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to retrieve user"})
	}
	return c.Status(200).JSON(fiber.Map{"data": user})
}

func CreateUsers(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	userCreated, err := services.CreateUsers(*user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to create users"})
	}
	return c.Status(200).JSON(fiber.Map{"data": userCreated, "message": "Success"})
}

func EditUsers(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	editedUser, err := services.EditUsers(*user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to edit user"})
	}

	return c.Status(200).JSON(fiber.Map{"data": editedUser, "message": "User updated successfully"})
}

func DeleteUser(c *fiber.Ctx) error {
	email := c.Params("email")
	err := services.DeleteUser(email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to delete user"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "User deleted successfully"})
}
