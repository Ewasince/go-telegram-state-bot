package context

import (
	"github.com/Ewasince/go-telegram-state-bot/api_utils"
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BaseBotContext struct {
	MessageText     string
	MessageCommand  string
	MessageSenderId int64
	MessageChatId   int64
	DefaultKeyboard Keyboard
	BotHandler      api_utils.SenderHandler
	CallCount       uint
	ErrorMessage    string
}

func NewContext(
	message *tg.Message,
	senderHandler *api_utils.BaseSenderHandler,
	errorMessage string,
) *BaseBotContext {
	return &BaseBotContext{
		MessageText:     message.Text,
		MessageCommand:  message.Command(),
		MessageSenderId: message.From.ID,
		MessageChatId:   message.Chat.ID,
		BotHandler:      senderHandler,
		ErrorMessage:    errorMessage,
	}
}

func (b *BaseBotContext) GetMessageCommand() string {
	return b.MessageCommand
}
func (b *BaseBotContext) GetMessageText() string {
	return b.MessageText
}
func (b *BaseBotContext) GetMessageSenderId() int64 {
	return b.MessageSenderId
}
func (b *BaseBotContext) GetMessageChatId() int64 { return b.MessageChatId }

func (b *BaseBotContext) SendChattables(keyboard Keyboard, stateChattables ...StateChattable) {
	for _, stateChattable := range stateChattables {
		if keyboard != nil {
			stateChattable.SetKeyboard(keyboard)
		} else {
			stateChattable.SetKeyboard(b.DefaultKeyboard)
		}
		if err := b.BotHandler.SendChattable(stateChattable.GetChattable()); err != nil {
			panic(err)
		}
	}
}

func (b *BaseBotContext) SetKeyboard(keyboard Keyboard) {
	b.DefaultKeyboard = keyboard
}

func (b *BaseBotContext) IncCallCount() uint {
	b.CallCount++
	return b.CallCount
}
