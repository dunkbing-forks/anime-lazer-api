package routes

import (
	"anime-lazer-api/services"
	"github.com/gofiber/fiber/v2"
)

func AnimeRoute(router fiber.Router) {
	animeRoute := router.Group("/anime")
	animeRoute.Get("/popular/:page", services.GetPopular)
}
