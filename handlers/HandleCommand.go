package handlers

import (
	"CheckVerifier/config"
	"CheckVerifier/db"
	"CheckVerifier/locales"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"time"
)

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
	case "users":
		userID := update.Message.From.ID
		InfoUsers(ctx, bot, queries, userID)
	default:
		_, text := locales.GetTranslation(ctx, bot, queries, "unknown_command", update)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		bot.Send(msg)
	}
}

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

func RegisterHandler(queries *db.Queries, bot *tgbotapi.BotAPI, update tgbotapi.Update, createUserFunc func(ctx context.Context, params db.CreateUserParams) error) {
	var userID int
	var username string
	var chatID int64
	if update.Message != nil {
		chatID = update.Message.Chat.ID
		userID = int(update.Message.From.ID)
		username = update.Message.From.UserName
	} else if update.CallbackQuery != nil {
		chatID = update.CallbackQuery.Message.Chat.ID
		userID = int(update.CallbackQuery.From.ID)
		username = update.CallbackQuery.From.UserName
	}
	params := db.CreateUserParams{
		UserID:   strconv.Itoa(userID),
		Username: username,
	}
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
	MessageOnStart(ctx, bot, queries, update, chatID)

	return
}

func MessageOnStart(ctx context.Context, bot *tgbotapi.BotAPI, queries *db.Queries, update tgbotapi.Update, chatID int64) {
	_, text := locales.GetTranslation(ctx, bot, queries, "start_message", update)
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)
	time.Sleep(1 * time.Second)

	_, text = locales.GetTranslation(ctx, bot, queries, "packet_information", update)
	msg = tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = locales.PacketKeyboard(bot, chatID, text)
	bot.Send(msg)
}

func InfoUsers(ctx context.Context, bot *tgbotapi.BotAPI, queries *db.Queries, chatID int64) {
	if isAdmin(chatID) {
		users, err := queries.SelectUsers(ctx)
		if err != nil {
			log.Fatalf("Failed to fetch users: %v", err)
		}

		var output string
		for _, user := range users {
			output += fmt.Sprintf(
				"ID: %d,  Username: %s, Purchased: %t,  Package: %s\n",
				user.ID, user.Username, user.Purchased, user.ChosenPackage,
			)
		}

		if output == "" {
			output = "No users found."
		}

		msg := tgbotapi.NewMessage(chatID, output)
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	} else {
		msg := tgbotapi.NewMessage(chatID, "denied\nyou are not an admin!")
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}

func isAdmin(chatID int64) bool {
	config.LoadConfig()
	return strconv.FormatInt(chatID, 10) == config.RECEIVER_ID
}
