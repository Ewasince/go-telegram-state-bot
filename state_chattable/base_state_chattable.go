package state_chattable

import (
	"github.com/Ewasince/go-telegram-state-bot/interfaces"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ interfaces.StateChattable = (*baseStateChattable)(nil) // interface hint

type baseStateChattable struct {
	chattable tg.Chattable
	baseChat  *tg.BaseChat
}

func NewBaseStateChattable(chattable tg.Chattable, baseChat *tg.BaseChat) interfaces.StateChattable {
	return &baseStateChattable{
		chattable: chattable,
		baseChat:  baseChat,
	}
}

func (b *baseStateChattable) GetChattable() tg.Chattable {
	return b.chattable
}
func (b *baseStateChattable) SetKeyboard(markup interfaces.Keyboard) {
	if markup != nil {
		b.baseChat.ReplyMarkup = markup.GetKeyBoard()
	} else {
		b.baseChat.ReplyMarkup = tg.NewRemoveKeyboard(true)
	}
}
