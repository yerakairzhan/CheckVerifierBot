package handlers

import (
	"CheckVerifier/db"
	"CheckVerifier/locales"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func RegisterHandler(queries *db.Queries, bot *tgbotapi.BotAPI, update tgbotapi.Update, createUserFunc func(ctx context.Context, params db.CreateUserParams) error) {
	//var chatID int64
	var userID int
	var username string

	// Determine the source of the request
	if update.Message != nil {
		//chatID = (update.Message.Chat.ID)
		userID = int(update.Message.From.ID)
		username = update.Message.From.UserName
	} else if update.CallbackQuery != nil {
		//chatID = (update.CallbackQuery.Message.Chat.ID)
		userID = int(update.CallbackQuery.From.ID)
		username = update.CallbackQuery.From.UserName
	}

	// Create the user
	params := db.CreateUserParams{
		UserID:   strconv.Itoa(userID),
		Username: username,
	}

	// Context for DB operation
	ctx := context.Background()

	err := createUserFunc(ctx, params)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return
	}

	if update.CallbackQuery != nil {
		callbackResponse := tgbotapi.NewCallback(update.CallbackQuery.ID, "Успешная регистрация.")
		bot.Request(callbackResponse)
	}

	text := locales.GetTranslation(ctx, bot, queries, "start_message", update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	bot.Send(msg)

	return
}
