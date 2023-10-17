package config

import (
	"os"

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
	godotenv.Load(".env")
	return Config{
		Port:         os.Getenv("PORT"),
		Host:         os.Getenv("HOST"),
		SMTPServer:   os.Getenv("SMTP_SERVER"),
		SMTPPort:     os.Getenv("SMTP_PORT"),
		SMTPUsername: os.Getenv("SMTP_USERNAME"),
		SMTPPassword: os.Getenv("SMTP_PASSWORD"),
	}, nil
}
