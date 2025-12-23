package routes

import (
	"fiber-api/internal/controllers"

	"github.com/gofiber/fiber/v3"
)

func UserRoute(api fiber.Router) {

	api.Get("/users", controllers.GetUsers)
	api.Get("/users/id/:id", controllers.GetUserById)
	api.Get("/users/email/:email", controllers.GetUserByEmail)
	api.Post("/users", controllers.CreateUsers)
	api.Put("/users/", controllers.EditUsers)
	api.Delete("/users/:email", controllers.DeleteUser)
	api.Post("/users/login", controllers.LoginUser)
}
