package utils

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	return logger
}
