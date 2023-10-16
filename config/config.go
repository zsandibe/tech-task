package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Host         string
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func NewConfig() (Config, error) {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		return Config{}, err
	}
	return Config{
		Port:         os.Getenv("PORT"),
		Host:         os.Getenv("HOST"),
		SMTPServer:   os.Getenv("SMTP_SERVER"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}, nil
}
