package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetRedisURI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return "redis://localhost:6379"
}

func GetMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGODB_URI")
}
