package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWT_SECRET string
}

var AppConfig Config

func Init() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	AppConfig = Config{
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	return nil
}
