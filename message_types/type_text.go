package message_types

import (
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
	"github.com/Ewasince/go-telegram-state-bot/state_chattable"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TextMessage string

//func NewTextMessage(messageText string) Messagables {
//	return TextMessage(messageText)
//}

func (t TextMessage) ToTgMessages(c BotContext) ([]StateChattable, error) {
	message := tg.NewMessage(c.GetMessageChatId(), string(t))
	stateChattable := state_chattable.NewBaseStateChattable(&message, &message.BaseChat)
	return []StateChattable{stateChattable}, nil
}
