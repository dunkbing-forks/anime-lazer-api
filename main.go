package main

import (
	"anime-lazer-api/configs"
	"anime-lazer-api/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	routes.UserRoute(app)
	err := app.Listen(":3000")
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
	}
}
