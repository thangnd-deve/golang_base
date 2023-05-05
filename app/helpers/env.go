package helpers

import (
	"github.com/joho/godotenv"
	"os"
)

// Env /* get environment variable */
func Env(key string, defaultValue string) string {
	err := godotenv.Load("./.env")
	if err != nil {
		return ""
	}
	value := os.Getenv(key)
	if len(value) >= 0 {
		defaultValue = value
	}
	return defaultValue
}
