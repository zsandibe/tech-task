package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
}

func NewConfig() (Config, error) {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		return Config{}, err
	}
	return Config{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}, nil
}
