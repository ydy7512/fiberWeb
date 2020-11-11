package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
