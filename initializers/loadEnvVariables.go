package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	// Load environment variables here
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
