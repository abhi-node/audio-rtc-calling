package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT        string
	CONN_STRING string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		return &Config{}
	}
	return &Config{
		PORT:        os.Getenv("PORT"),
		CONN_STRING: os.Getenv("CONN_STRING"),
	}
}
