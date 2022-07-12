package routes

import (
	"anime-lazer-api/services"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(router fiber.Router) {
	router.Post("/users", services.CreateUser)
	router.Get("/users", services.GetAllUsers)
	router.Get("/users/:userId", services.GetAUser)
	router.Put("/users/:userId", services.UpdateAUser)
	router.Delete("/users/:userId", services.DeleteAUser)
}
