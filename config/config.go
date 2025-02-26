package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port   string
	Env    string
	DbHost string
	DbName string
	DbUser string
	DbPass string
	DbPort string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Port:   os.Getenv("PORT"),
		Env:    os.Getenv("ENV"),
		DbHost: os.Getenv("DB_HOST"),
		DbName: os.Getenv("DB_NAME"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbPort: os.Getenv("DB_PORT"),
	}
}

const DEV_ENV = "DEV"
const STAGE_ENV = "STAGE"
const PROD_ENV = "PROD"
