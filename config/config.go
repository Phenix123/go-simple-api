package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
	Env  string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),
	}
}

const DEV_ENV = "DEV"
const STAGE_ENV = "STAGE"
const PROD_ENV = "PROD"
