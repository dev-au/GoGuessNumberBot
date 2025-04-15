package handlers

import (
	"GoTestBot/config"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	currentState := config.GetCurrentState(update.Message.Chat.ID)
	if currentState != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "O'yin to'xtatildi. Qayta boshlash uchun /go ni bosing")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		config.ClearState(update.Message.Chat.ID)
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Salom, bu son topish o'yini, 1dan 100gacha son o'yla men uni ko'pi bilan 7 ta urinishda topaman."+
			" Boshlash uchun /go ni bosing",
	)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)

}

func StartGame(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	firstNumber := 50
	var minMaxKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("< Katta"),
			tgbotapi.NewKeyboardButton("> Kichik"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Teng"),
		),
	)
	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		fmt.Sprintf("Unda o'yinni boshladik. Sen o'ylagan son %d mi? Yoki bu son", firstNumber),
	)
	msg.ReplyMarkup = minMaxKeyboard
	bot.Send(msg)
	config.SetState(update.Message.Chat.ID, "guess")
	config.SetStateData(update.Message.Chat.ID, "max", 100)
	config.SetStateData(update.Message.Chat.ID, "min", 1)
	config.SetStateData(update.Message.Chat.ID, "try", firstNumber)
	config.SetStateData(update.Message.Chat.ID, "count", 1)
}

func GameProcess(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	currentData := config.GetStateData(update.Message.Chat.ID)
	minNumber := currentData["min"].(int)
	maxNumber := currentData["max"].(int)
	currentGuess := currentData["try"].(int)
	guessCount := currentData["count"].(int)
	newHint := update.Message.Text
	var newGuess int
	if newHint == "< Katta" {
		newGuess = (minNumber + currentGuess) / 2
		config.SetStateData(update.Message.Chat.ID, "max", currentGuess)
		config.SetStateData(update.Message.Chat.ID, "min", minNumber)
	} else if newHint == "> Kichik" {

		if currentGuess == maxNumber {
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Aldama men uni 7 urinishda topib bo'lgan edim. Qayta o'yin uchun /go ni bosing",
			)
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			config.ClearState(update.Message.Chat.ID)
			return
		}
		newGuess = (currentGuess + maxNumber) / 2
		config.SetStateData(update.Message.Chat.ID, "max", maxNumber)
		config.SetStateData(update.Message.Chat.ID, "min", currentGuess)
	} else {
		msg := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Ajoyib men sen o'ylagan sonni %d ta urinishda topdim. Qayta o'yin uchun /go ni bosing",
				guessCount),
		)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
		config.ClearState(update.Message.Chat.ID)
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		fmt.Sprintf("Xo'p unda %d bu sen o'ylagan sonmi", newGuess),
	)
	config.SetStateData(update.Message.Chat.ID, "try", newGuess)
	config.SetStateData(update.Message.Chat.ID, "count", guessCount+1)
	bot.Send(msg)
}
