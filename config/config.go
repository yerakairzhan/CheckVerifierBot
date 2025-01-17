package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var BotToken string
var DB_HOST string
var DB_PORT string
var DB_USER string
var DB_PASSWORD string
var DB_NAME string
var PAY_URL string
var RECEIVER_ID string

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
	PAY_URL = os.Getenv("PAY_URL")
	RECEIVER_ID = os.Getenv("USER_CHECK_ID")
}
