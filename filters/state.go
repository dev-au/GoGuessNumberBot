package filters

import (
	"GoTestBot/config"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StateFilter(stateKey string, filterFunctions ...func(update tgbotapi.Update) bool) func(update tgbotapi.Update) bool {
	return func(update tgbotapi.Update) bool {
		var chatID int64
		if update.Message != nil {
			chatID = update.Message.Chat.ID
		} else if update.CallbackQuery != nil {
			chatID = update.CallbackQuery.Message.Chat.ID
		} else {
			fmt.Println("Another Type")
			return false
		}
		if config.GetCurrentState(chatID) == stateKey {
			return applyFilters(update, filterFunctions...)
		}
		return false
	}

}
