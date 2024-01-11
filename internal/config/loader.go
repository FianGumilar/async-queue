package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return &Config{
		Redis: Redis{
			Addr: os.Getenv("REDIS_ADDR"),
			Pass: os.Getenv("REDIS_PASS"),
		},
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Email: Email{
			Host: os.Getenv("EMAIL_HOST"),
			Port: os.Getenv("EMAIL_PORT"),
			User: os.Getenv("EMAIL_USER"),
			Pass: os.Getenv("EMAIL_PASS"),
		},
	}
}
