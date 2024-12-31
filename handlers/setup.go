package handlers

import (
	"CheckVerifier/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupHandlers(bot *tgbotapi.BotAPI, queries *db.Queries) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			handleCallback(bot, update, queries)
		} else if update.Message != nil {
			chatID := update.Message.Chat.ID
			userID := update.Message.From.ID
			if update.Message.IsCommand() {
				handleCommand(bot, update, queries)
			} else if update.Message.Photo != nil {
				handlePhoto(bot, update, queries, userID)
			} else {
				handleReply(bot, update, queries, chatID, userID)
			}
		}
	}
}
