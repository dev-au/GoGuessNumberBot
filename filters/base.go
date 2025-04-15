package filters

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func applyFilters(update tgbotapi.Update, filterFunctions ...func(update tgbotapi.Update) bool) bool {
	for _, filterFunc := range filterFunctions {
		if filterFunc != nil && !filterFunc(update) {
			return false
		}
	}
	return true
}
