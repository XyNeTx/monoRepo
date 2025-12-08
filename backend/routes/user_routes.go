package routes

import (
	"fiber-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router) {

	api := app.Group("/api")

	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:email", controllers.GetUserByEmail)
	api.Post("/users", controllers.CreateUsers)
	api.Put("/users/", controllers.EditUsers)
	api.Delete("/users/:email", controllers.DeleteUser)
}
