package interfaces

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Keyboard interface {
	GetKeyBoard() *tg.ReplyKeyboardMarkup
	ProcessMessage(BotContext) (HandlerResponse, bool)
}

//func (b *BotKeyboard) GetKeyBoard() tg.ReplyKeyboardMarkup {
//
//	func (b *BotKeyboard) ProcessMessage(c BotContext) (HandlerResponse, bool) {
