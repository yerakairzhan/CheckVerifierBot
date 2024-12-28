package handlers

import (
	"CheckVerifier/db"
	"CheckVerifier/locales"
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetupHandlers(bot *tgbotapi.BotAPI, queries *db.Queries) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go func(update tgbotapi.Update) {
			if update.CallbackQuery != nil {
				log.Printf("Received CallbackQuery: %+v", update.CallbackQuery)
				HandleCallback(bot, update, queries)
			} else if update.Message != nil && update.Message.IsCommand() {
				handleCommand(bot, update, queries)
			} else {
				log.Printf("Unhandled update: %+v\n", update)
			}
		}(update)
	}
}

func handleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries) {
	ctx := context.Background()
	switch update.Message.Command() {
	case "start":
		log.Printf("Started by: @%s", update.Message.From.UserName)
		RegisterHandler(queries, bot, update, func(ctx context.Context, params db.CreateUserParams) error {
			return queries.CreateUser(ctx, params)
		})
	case "lang":
		AskForLanguage(queries, bot, update)
		log.Printf("Language handler called by: @%s", update.Message.From.UserName)

	default:
		text := locales.GetTranslation(ctx, bot, queries, "unknown_command", update)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
	}
}
