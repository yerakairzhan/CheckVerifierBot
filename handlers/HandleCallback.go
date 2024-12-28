package handlers

import (
	"CheckVerifier/db"
	"context"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries) {
	callbackData := update.CallbackQuery.Data
	chatID := update.CallbackQuery.Message.Chat.ID
	userID := update.CallbackQuery.From.ID
	log.Printf("Received callback: %s from user %d", callbackData, userID)

	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Language selection received.")
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Failed to acknowledge callback: %v", err)
	}

	switch callbackData {
	case "callback_data_eng":
		log.Println("Setting language to English for user:", userID)
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "en",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		msg := tgbotapi.NewMessage(chatID, "Language changed to English.")
		bot.Send(msg)
	case "callback_data_kaz":
		log.Println("Setting language to Kazakh for user:", userID)
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "kz",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		msg := tgbotapi.NewMessage(chatID, "Language changed to Kazakh.")
		bot.Send(msg)
	case "callback_data_rus":
		log.Println("Setting language to Russian for user:", userID)
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "ru",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		msg := tgbotapi.NewMessage(chatID, "Language changed to Russian.")
		bot.Send(msg)
	default:
		log.Printf("Unknown callback data: %s", callbackData)
		msg := tgbotapi.NewMessage(chatID, "Unknown action.")
		bot.Send(msg)
	}
}
