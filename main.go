package main

import (
	"CheckVerifier/config"
	"CheckVerifier/db"
	"CheckVerifier/handlers"
	"CheckVerifier/locales"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	// Load translations at startup
	err := locales.LoadTranslations()
	if err != nil {
		log.Fatalf("Failed to load translations: %v", err)
	}

	config.LoadConfig()

	// Загрузка Базы данных с параметров конфигурации
	envPost := config.DB_HOST
	envPort := config.DB_PORT
	envUser := config.DB_USER
	envPass := config.DB_PASSWORD
	envDbnm := config.DB_NAME

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envPost, envPort, envDbnm)
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error with DB connection: %v", err)
	}
	defer dbConn.Close()
	// Проверка соединения с базой данных
	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	// Инициализация SQLC Queries
	queries := db.New(dbConn)
	// Настройка маршрутов команд
	handlers.SetupHandlers(bot, queries)
}
