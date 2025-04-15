package filters

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func TextFilter(text string, filterFunctions ...func(update tgbotapi.Update) bool) func(update tgbotapi.Update) bool {
	return func(update tgbotapi.Update) bool {
		if update.Message == nil || update.Message.Text != text {
			return false
		}
		return applyFilters(update, filterFunctions...)
	}
}

func NoMessageFilter(filterFunctions ...func(update tgbotapi.Update) bool) func(update tgbotapi.Update) bool {
	return func(update tgbotapi.Update) bool {
		return applyFilters(update, filterFunctions...)
	}
}
