package services

import (
	"anime-lazer-api/configs"
	"anime-lazer-api/responses"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"time"
)

var animeCollection = configs.GetCollection(configs.DB, "animes")

func GetPopular(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	page, _ := strconv.ParseInt(c.Params("page"), 10, 64)
	findOptions := options.Find()
	Paginate(page, 10, findOptions)
	result, err := animeCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}
	var animes []map[string]interface{}
	for result.Next(ctx) {
		//var singleUser models.User
		var anime map[string]interface{}
		if err = result.Decode(&anime); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
		animes = append(animes, anime)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": animes}},
	)
}
