package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		return &Config{}
	}
	return &Config{
		PORT: os.Getenv("PORT"),
	}
}
