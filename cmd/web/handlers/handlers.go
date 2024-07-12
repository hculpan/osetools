package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var appTitle string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("missing .env file: %s", err.Error())
	}

	appTitle = os.Getenv("APP_TITLE")
	if appTitle == "" {
		log.Fatal("missing APP_TITLE in configuration file")
	}
}
