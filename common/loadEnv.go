package common

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading file env %s", err)
	}
}
