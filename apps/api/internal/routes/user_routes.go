package routes

import (
	"fiber-api/internal/controllers"

	"github.com/gofiber/fiber/v3"
)

func UserRoute(api fiber.Router, uc *controllers.UserController) {

	api.Get("/users", uc.GetUsers)
	api.Get("/users/id/:id", uc.GetUserById)
	api.Get("/users/email/:email", uc.GetUserByEmail)
	api.Post("/users/signup", uc.CreateUsers)
	api.Put("/users/", uc.EditUsers)
	api.Delete("/users/:email", uc.DeleteUser)
	api.Post("/users/login", uc.LoginUser)
	api.Patch("/users/change-password", uc.ChangePassword)
}
