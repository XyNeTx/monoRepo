package routes

import (
	"fiber-api/controllers"
	"fiber-api/database"
	"fiber-api/models"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router) {
	database.ConnectDB()

	database.DB.AutoMigrate(&models.User{})

	api := app.Group("/api")

	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:email", controllers.GetUserByEmail)
	api.Post("/users", controllers.CreateUsers)
	api.Put("/users/", controllers.EditUsers)
	api.Delete("/users/:email", controllers.DeleteUser)
}
