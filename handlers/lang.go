package handlers

import (
	"CheckVerifier/db"
	"CheckVerifier/locales"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func AskForLanguage(queries *db.Queries, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var chatID int64
	ctx := context.Background()
	_, text := locales.GetTranslation(ctx, bot, queries, "change_lang_message", update)
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	} else if update.CallbackQuery != nil {
		chatID = update.Message.Chat.ID
	}
	log.Print("AskForLanguage ended")
	locales.InlineLanguage(bot, chatID, text)
}
