package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var BotToken string
var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var DB_NAME string

func LoadConfig() {
	// Загрузка переменных из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	BotToken = os.Getenv("BOT_TOKEN")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
}
