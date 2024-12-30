package handlers

import (
	"CheckVerifier/db"
	"CheckVerifier/locales"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries) {
	callbackData := update.CallbackQuery.Data
	chatID := update.CallbackQuery.Message.Chat.ID
	userID := update.CallbackQuery.From.ID
	messageID := update.CallbackQuery.Message.MessageID
	ctx := context.Background()
	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "Language selection received.")
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Failed to acknowledge callback: %v", err)
	}

	switch callbackData {
	case "callback_data_eng":
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "en",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		saySuccessful(ctx, bot, queries, "success_language_change", update, messageID)
	case "callback_data_kaz":
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "kz",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		saySuccessful(ctx, bot, queries, "success_language_change", update, messageID)
	case "callback_data_rus":
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "ru",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		saySuccessful(ctx, bot, queries, "success_language_change", update, messageID)

	default:
		log.Printf("Unknown callback data: %s", callbackData)
		msg := tgbotapi.NewMessage(chatID, "Unknown action.")
		bot.Send(msg)
	}
	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	if _, err := bot.Request(deleteMsg); err != nil {
		log.Printf("Failed to delete message: %v", err)
	}
}

func handleReply(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries, chatID int64) {
	userResponse := update.Message.Text
	ctx := context.Background()
	_, text := locales.GetTranslation(ctx, bot, queries, "follow", update)
	removeKeyboardMsg := tgbotapi.NewMessage(chatID, text)
	removeKeyboardMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(removeKeyboardMsg)
	switch userResponse {
	case "NameOfFirst":
		_, text := locales.GetTranslation(ctx, bot, queries, "packet_1", update)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = locales.LinkKeyboard()
		bot.Send(msg)
	case "NameOfSecond":
		_, text := locales.GetTranslation(ctx, bot, queries, "packet_2", update)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = locales.LinkKeyboard()
		bot.Send(msg)
	case "NameOfThird":
		_, text := locales.GetTranslation(ctx, bot, queries, "packet_3", update)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = locales.LinkKeyboard()
		bot.Send(msg)
	default:
		var msg tgbotapi.MessageConfig
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	}
}

func saySuccessful(ctx context.Context, bot *tgbotapi.BotAPI, queries *db.Queries, key string, update tgbotapi.Update, messageID int) {
	chatID := update.CallbackQuery.Message.Chat.ID
	err, text := locales.GetTranslation(ctx, bot, queries, "success_language_change", update)
	if err != nil {
		log.Printf("Failed to get translation: %v", err)
	}
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Failed to send message: %v", err)
	}

	MessageOnStart(ctx, bot, queries, update, chatID)
}
