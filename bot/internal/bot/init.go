package bot

import tgbotapi "github.com/Syfaro/telegram-bot-api"

func New() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI("token") // os.Getenv("TOKEN")
	if err != nil {
		return nil, err
	}
	bot.Debug = true // os.Getenv("LOG_LEVEL")

	tgbotapi.NewInlineKeyboardMarkup()

	return bot, nil
}
