package main

import (
	"fiber-api/routes"

	// นำเข้าแพ็คเกจ Gin
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New() // สร้าง router ด้วย Gin

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "Welcome to the Fiber API"})
	})

	routes.UserRoute(app)

	app.Listen(":8080") // รันเซิร์ฟเวอร์บน port 8080
}
