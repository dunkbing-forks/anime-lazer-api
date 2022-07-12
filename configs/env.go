package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv("MONGO_URI")
}
