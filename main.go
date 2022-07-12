package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	err := app.Listen(":3000")
	if err != nil {
		_ = fmt.Errorf("error: %v", err)
	}
}