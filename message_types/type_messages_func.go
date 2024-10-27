package message_types

import (
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
)

var _ Messagables = (*BotMessageHandler)(nil) // interface hint

type BotMessageHandler func(c BotContext) (Messagables, error)

func (b BotMessageHandler) ToTgMessages(c BotContext) ([]StateChattable, error) {
	res, err := b(c)
	if err != nil {
		return nil, err
	}
	return res.ToTgMessages(c)
}
