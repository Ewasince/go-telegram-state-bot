package teleBotStateLib

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type BotButton struct {
	ButtonTitle   string
	ButtonHandler ContextHandler
}

type ButtonsRow []BotButton
type BotKeyboard struct {
	Keyboard []ButtonsRow
}

func (b *BotKeyboard) GetKeyBoard() tg.ReplyKeyboardMarkup {
	var buttonsArray [][]tg.KeyboardButton

	for _, row := range b.Keyboard {
		var buttonsRow []tg.KeyboardButton
		for _, button := range row {
			buttonsRow = append(buttonsRow, tg.KeyboardButton{
				Text: button.ButtonTitle,
			})
		}
		buttonsArray = append(buttonsArray, buttonsRow)
	}

	keyboard := tg.ReplyKeyboardMarkup{
		Keyboard: buttonsArray,
	}
	return keyboard
}

// ProcessMessage return bot state id, is new state, is button pressed and error
func (b *BotKeyboard) ProcessMessage(c BotContext) (HandlerResponse, bool) {
	for _, row := range b.Keyboard {
		for _, button := range row {
			if button.ButtonTitle == c.GetMessageText() {
				handlerResponse := button.ButtonHandler(c)
				return handlerResponse, true
			}
		}
	}
	return HandlerResponse{}, false
}
