package helpers

import (
	"github.com/Ewasince/go-telegram-state-bot/interfaces"
	"github.com/Ewasince/go-telegram-state-bot/message_types"
)

func CreateAndSendMessage(messageText string, c interfaces.BotContext) {
	textMessages, err := message_types.TextMessage(messageText).ToTgMessages(c)
	if err != nil {
		panic(err)
	}
	c.SendChattables(nil, textMessages[0])
}
