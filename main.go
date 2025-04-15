package main

import (
	"GoTestBot/config"
	"GoTestBot/filters"
	"GoTestBot/handlers"
)

var botHandlers = []config.BotHandler{
	{filters.TextFilter("/start"), handlers.StartCommand},
	{filters.TextFilter("/go"), handlers.StartGame},
	{filters.NoMessageFilter(
		filters.StateFilter("guess"),
	), handlers.GameProcess},
}

func main() {
	config.FetchUpdates(botHandlers)

}
