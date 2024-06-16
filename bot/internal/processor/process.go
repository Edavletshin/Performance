package processor

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	// TODO: add desc
	commandStart     = "/start"
	commandStartDesc = "start"

	commandCreate     = "/create"
	commandCreateDesc = "create"

	commandDonate     = "/donate"
	commandDonateDesc = "donate"

	commandShowStatus     = "/show_status"
	commandShowStatusDesc = "show_status"

	defaultDesc = "Unexpected command, please use one of allowed:  " // TODO: all commands desc??
)

var (
	mainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(commandCreate, commandCreateDesc),
			tgbotapi.NewInlineKeyboardButtonData(commandDonate, commandDonateDesc),
			tgbotapi.NewInlineKeyboardButtonData(commandShowStatus, commandShowStatusDesc),
			// tgbotapi.NewInlineKeyboardButtonSwitch()
		),
	)
)

func Process(bot *tgbotapi.BotAPI) error {
	// Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0) // os.Getenv("UPDATE_OFFSET")
	u.Timeout = 60             // os.Getenv("UPDATE_TIMEOUT")

	// Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//// Construct a new message from the given chat ID and containing
		//// the text that we received.
		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//
		//// If the message was open, add a copy of our numeric keyboard.
		//switch update.Message.Text {
		//case commandStart:
		//	msg.ReplyMarkup = mainKeyboard
		//	msg.Text = commandStartDesc
		//}
		//
		//// Send the message.
		//if _, err = bot.Send(msg); err != nil {
		//	fmt.Errorf("%v", err)
		//}

		////Проверяем что от пользователья пришло именно текстовое сообщение
		//if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
		//	continue
		//}

		switch update.Message.Text {
		case commandStart:
			botSendMessage(bot, update.Message.Chat.ID, mainKeyboard, commandStartDesc)

		case commandCreate:
			botSendMessage(bot, update.Message.Chat.ID, mainKeyboard, commandCreateDesc)

		case commandDonate:
			botSendMessage(bot, update.Message.Chat.ID, mainKeyboard, commandDonateDesc)

		case commandShowStatus:
			botSendMessage(bot, update.Message.Chat.ID, mainKeyboard, commandShowStatusDesc)

		default:
			botSendMessage(bot, update.Message.Chat.ID, mainKeyboard, defaultDesc)
		}
	}
	return nil
}

func botSendMessage(bot *tgbotapi.BotAPI, chatID int64, keyboard tgbotapi.InlineKeyboardMarkup, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyMarkup = keyboard
	_, err := bot.Send(msg)
	if err != nil {
		fmt.Errorf("while sending: %v", err)
	}
}
