package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTPPort string

	RESERVATION_SERVICE_PORT string
	PAYMENT_SERVICE_PORT     string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPPort = cast.ToString(coalesce("HTTP_PORT", ":8080"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))
	config.RESERVATION_SERVICE_PORT = cast.ToString(coalesce("RESERVATION_SERVICE_PORT", ":50051"))
	config.PAYMENT_SERVICE_PORT = cast.ToString(coalesce("PAYMENT_SERVICE_PORT", ":50052"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
