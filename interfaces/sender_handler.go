package interfaces

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type SenderHandler interface {
	SendChattable(tg.Chattable) error
}
