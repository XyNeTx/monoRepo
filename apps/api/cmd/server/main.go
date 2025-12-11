package main

import (
	"fiber-api/internal/config"
	"fiber-api/internal/database"
	"fiber-api/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/swagger/v2"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()
	database.Migrate()

	app := fiber.New() // สร้าง router ด้วย

	// app.Use(func(c *fiber.Ctx) error {
	// 	c.Set("Access-Control-Allow-Origin", "*")
	// 	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	// 	if c.Method() == "OPTIONS" {
	// 		return c.SendStatus(204)
	// 	}

	// 	return c.Next()
	// })
	app.Use(cors.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Welcome to the Fiber API"})
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	api := app.Group("/api")
	routes.UserRoute(api)

	app.Listen(":8080") // รันเซิร์ฟเวอร์บน port 8080
}
