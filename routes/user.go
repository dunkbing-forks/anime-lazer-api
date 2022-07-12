package routes

import (
	"anime-lazer-api/services"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/users", services.GetAllUsers)
	app.Post("/users", services.CreateUser)
	app.Get("/users/:userId", services.GetAUser)
	app.Put("/users/:userId", services.UpdateAUser)
	app.Delete("/users/:userId", services.DeleteAUser)
}
