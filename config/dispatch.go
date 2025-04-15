package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandlerFunc func(bot *tgbotapi.BotAPI, update tgbotapi.Update)
type FilterFunc func(update tgbotapi.Update) bool

type BotHandler struct {
	Filter  func(update tgbotapi.Update) bool
	Handler func(bot *tgbotapi.BotAPI, update tgbotapi.Update)
}

func DispatchBotHandler(filter FilterFunc, handler BotHandlerFunc) func(bot *tgbotapi.BotAPI, update tgbotapi.Update) bool {
	return func(bot *tgbotapi.BotAPI, update tgbotapi.Update) bool {
		if filter(update) {
			handler(bot, update)
			return true
		}
		return false
	}
}
