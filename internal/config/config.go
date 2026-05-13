package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() Config {

	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("PORT not set, defaulting to 8080")
	}

	return Config{
		Port: port,
	}
}
