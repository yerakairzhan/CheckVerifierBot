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
)

func handlePhoto(bot *tgbotapi.BotAPI, update tgbotapi.Update, queries *db.Queries, userID int64) {
	ctx := context.Background()
	config.LoadConfig()
	receiver := config.RECEIVER_ID

	if update.Message.Photo == nil || len(update.Message.Photo) == 0 {
		log.Println("No photo found in message.")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No photo detected. Please try again.")
		bot.Send(msg)
		return
	}

	photo := update.Message.Photo[len(update.Message.Photo)-1] // Last photo is the highest resolution

	RecieverID, err := strconv.ParseInt(receiver, 10, 64)
	if err != nil {
		log.Println(err)
	}

	info, err := queries.InfoUser(ctx, strconv.FormatInt(userID, 10))
	if err != nil {
		log.Fatalf("Failed to retrieve user info: %v", err)
	}

	// Assert the type to string
	formattedOutput, ok := info.(string)
	if !ok {
		log.Fatalf("Unexpected type for formattedOutput: %T", info)
	}

	photoMsg := tgbotapi.NewPhoto(RecieverID, tgbotapi.FileID(photo.FileID))
	photoMsg.Caption = fmt.Sprintf(formattedOutput)
	photoMsg.ReplyMarkup = locales.InlineForAdmin(strconv.FormatInt(userID, 10))

	if _, err := bot.Send(photoMsg); err != nil {
		log.Printf("Failed to send photo to receiver: %v", err)
		return
	}

	_, text := locales.GetTranslation(ctx, bot, queries, "sended_photo", update)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Failed to send confirmation message: %v", err)
	}

}
