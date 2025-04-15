package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func FetchUpdates(handlers []BotHandler) {
	bot, u := SetupBot()
	updates := bot.GetUpdatesChan(u)
	var configuredHandlers []func(bot *tgbotapi.BotAPI, update tgbotapi.Update) bool

	for _, handler := range handlers {
		configuredHandlers = append(configuredHandlers, DispatchBotHandler(handler.Filter, handler.Handler))
	}

	for update := range updates {
		log.Printf("[%s] Received new update", bot.Self.UserName)

		for _, handler := range configuredHandlers {

			if handler(bot, update) {
				break
			}
		}
	}
}
