package teleBotStateLib

import (
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
)

type BotCommand struct {
	CommandMessage string
	CommandHandler ContextHandler
}
