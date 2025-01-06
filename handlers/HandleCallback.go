package handlers

import (
	"CheckVerifier/db"
	"CheckVerifier/locales"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func handleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries) {
	callbackData := update.CallbackQuery.Data
	chatID := update.CallbackQuery.Message.Chat.ID
	userID := update.CallbackQuery.From.ID
	messageID := update.CallbackQuery.Message.MessageID
	ctx := context.Background()

	if callbackData == "callback_data_eng" {
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "en",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		saySuccessful(ctx, bot, queries, "success_language_change", update, messageID)
		deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
		if _, err := bot.Request(deleteMsg); err != nil {
			log.Printf("Failed to delete message: %v", err)
		}
	} else if callbackData == "callback_data_kaz" {
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "kz",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		saySuccessful(ctx, bot, queries, "success_language_change", update, messageID)
		deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
		if _, err := bot.Request(deleteMsg); err != nil {
			log.Printf("Failed to delete message: %v", err)
		}
	} else if callbackData == "callback_data_rus" {
		err := queries.ChangeLanguage(context.Background(), db.ChangeLanguageParams{
			UserID:       strconv.FormatInt(userID, 10),
			LanguageCode: "ru",
		})
		if err != nil {
			log.Printf("Failed to change language: %v", err)
		}
		saySuccessful(ctx, bot, queries, "success_language_change", update, messageID)
		deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
		if _, err := bot.Request(deleteMsg); err != nil {
			log.Printf("Failed to delete message: %v", err)
		}
	} else if strings.HasPrefix(callbackData, "accept_") {
		RecieverID, err := strconv.ParseInt(strings.TrimPrefix(callbackData, "accept_"), 10, 64)
		if err != nil {
			fmt.Printf("Error converting string to int64: %v\n", err)
			return
		}

		err = queries.AcceptPurchase(ctx, strconv.FormatInt(RecieverID, 10))
		if err != nil {
			return
		}
		_, text := locales.GetTranslation(ctx, bot, queries, "purchase_successful", update)
		msg := tgbotapi.NewMessage(RecieverID, text)
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send confirmation message: %v", err)
		}
	} else if strings.HasPrefix(callbackData, "reject_") {
		RecieverID, err := strconv.ParseInt(strings.TrimPrefix(callbackData, "reject_"), 10, 64)
		if err != nil {
			fmt.Printf("Error converting string to int64: %v\n", err)
			return
		}

		_, text := locales.GetTranslation(ctx, bot, queries, "purchase_unsuccessful", update)
		msg := tgbotapi.NewMessage(RecieverID, text)
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send confirmation message: %v", err)
		}
	} else if callbackData == "first_choosen" {
		_, text := locales.GetTranslation(ctx, bot, queries, "packet_1", update)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = locales.LinkKeyboard()
		bot.Send(msg)
		_, text = locales.GetTranslation(ctx, bot, queries, "packet_1_name", update)
		err := queries.SetPackage(context.Background(), db.SetPackageParams{
			UserID:        strconv.FormatInt(userID, 10),
			ChosenPackage: text,
		})
		if err != nil {
			log.Printf("Failed to set package: %v", err)
		}
	} else if callbackData == "first_choosen" {
		_, text := locales.GetTranslation(ctx, bot, queries, "packet_2", update)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = locales.LinkKeyboard()
		bot.Send(msg)
		_, text = locales.GetTranslation(ctx, bot, queries, "packet_2_name", update)
		err := queries.SetPackage(context.Background(), db.SetPackageParams{
			UserID:        strconv.FormatInt(userID, 10),
			ChosenPackage: text,
		})
		if err != nil {
			log.Printf("Failed to set package: %v", err)
		}
	} else if callbackData == "first_choosen" {
		_, text := locales.GetTranslation(ctx, bot, queries, "packet_3", update)
		msg := tgbotapi.NewMessage(chatID, text)
		msg.ReplyMarkup = locales.LinkKeyboard()
		bot.Send(msg)
		_, text = locales.GetTranslation(ctx, bot, queries, "packet_3_name", update)
		err := queries.SetPackage(context.Background(), db.SetPackageParams{
			UserID:        strconv.FormatInt(userID, 10),
			ChosenPackage: text,
		})
		if err != nil {
			log.Printf("Failed to set package: %v", err)
		}
	} else {
		log.Printf("Unknown callback data: %s", callbackData)
		_, text := locales.GetTranslation(ctx, bot, queries, "unknown_command", update)
		msg := tgbotapi.NewMessage(chatID, text)
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
