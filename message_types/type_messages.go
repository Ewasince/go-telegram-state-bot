package message_types

import (
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
)

var _ Messagables = (*BotMessages)(nil) // interface hint

type BotMessages []Messagables

func (b BotMessages) ToTgMessages(c BotContext) ([]StateChattable, error) {
	var tgChattables []StateChattable
	for _, messagable := range b {
		stateChattables, err := messagable.ToTgMessages(c)
		if err != nil {
			continue // FIXME: надо понять чё делать то
		}
		tgChattables = append(tgChattables, stateChattables...)
	}
	return tgChattables, nil
}
