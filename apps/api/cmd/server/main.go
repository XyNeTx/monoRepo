package main

import (
	"fiber-api/internal/config"
	"fiber-api/internal/controllers"
	"fiber-api/internal/database"
	_ "fiber-api/internal/docs"
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

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "doc.json",
	}))

	app.Get("/", controllers.InitialUserController)

	api := app.Group("/api")
	routes.UserRoute(api)

	app.Listen(":8080") // รันเซิร์ฟเวอร์บน port 8080
}
