package config

import (
	"github.com/joho/godotenv"
)

// LoadENV - load env file.
func LoadENV() {
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found")
	}
}
