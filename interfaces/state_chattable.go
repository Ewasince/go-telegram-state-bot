package interfaces

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type StateChattable interface {
	GetChattable() tg.Chattable
	SetKeyboard(keyboard Keyboard)
}
