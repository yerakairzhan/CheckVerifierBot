package locales

import (
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
	msg.ReplyMarkup = inlineKeyboard

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Failed to send inline keyboard: %v", err)
	}
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
