package config

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	JWT_SECRET        string
	TOKEN_COOKIE_NAME string
	TOKEN_COOKIE_AGE  int // in minutes
	BCRYPT_COST       int
}

var AppConfig Config

func Init() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	AppConfig = Config{
		JWT_SECRET:        os.Getenv("JWT_SECRET"),
		TOKEN_COOKIE_NAME: "token",
		TOKEN_COOKIE_AGE:  60,
		BCRYPT_COST:       bcrypt.DefaultCost,
	}

	return nil
}
