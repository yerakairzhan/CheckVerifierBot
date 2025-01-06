package locales

import (
	"CheckVerifier/config"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func InlineLanguage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("English", "callback_data_eng"),
			tgbotapi.NewInlineKeyboardButtonData("Kazakh", "callback_data_kaz"),
			tgbotapi.NewInlineKeyboardButtonData("Russian", "callback_data_rus"),
		),
	)
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	msg.ReplyMarkup = inlineKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Failed to send inline keyboard: %v", err)
	}
}

func LinkKeyboard() tgbotapi.InlineKeyboardMarkup {
	config.LoadConfig()
	url := config.PAY_URL
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("Kaspi tap", url)),
	)
	return inlineKeyboard
}

func InlineForAdmin(userID string) tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Accept", fmt.Sprintf("accept_%s", userID)),
			tgbotapi.NewInlineKeyboardButtonData("Reject", fmt.Sprintf("reject_%s", userID)),
		),
	)
	return inlineKeyboard
}

func PacketKeyboard(bot *tgbotapi.BotAPI, chatID int64, text string) tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("NameOfFirst"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("NameOfSecond"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("NameOfThird"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Отмена"),
		),
	)
	return keyboard
}

func InlinePacketKeyboard(first string, second string, third string) tgbotapi.InlineKeyboardMarkup {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(first, "first_choosen"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(second, "second_choosen"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(third, "third_choosen"),
		),
	)
	return inlineKeyboard
}
