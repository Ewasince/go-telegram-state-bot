package teleBotStateLib

import (
	"github.com/Ewasince/go-telegram-state-bot/apiUtils"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotContext interface {
	GetMessageCommand() string
	GetMessageText() string
	GetMessageSenderId() int64

	SendMessages(...tg.Chattable)
	CreateMessages(...string) []tg.MessageConfig

	CreateAndSendMessage(string)

	SetKeyboard(*BotKeyboard)

	SendErrorMessage()
	incCallCount() uint
}

type BaseBotContext struct {
	MessageText     string
	MessageCommand  string
	MessageSenderId int64
	MessageChatId   int64
	DefaultKeyboard *BotKeyboard
	BotHandler      apiUtils.SenderHandler
	CallCount       uint
	ErrorMessage    string
}

func NewContext(
	message *tg.Message,
	senderHandler *apiUtils.BaseSenderHandler,
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

func (b *BaseBotContext) SendMessages(chattables ...tg.Chattable) {
	for _, msg := range chattables {
		if err := b.BotHandler.SendMessage(msg); err != nil {
			panic(err)
		}
	}
}

func (b *BaseBotContext) CreateMessages(messages ...string) []tg.MessageConfig {
	var chattableMessages []tg.MessageConfig
	for _, msg := range messages {
		message := tg.NewMessage(b.MessageChatId, msg)
		if b.DefaultKeyboard != nil {
			message.ReplyMarkup = b.DefaultKeyboard.GetKeyBoard()
		} else {
			message.ReplyMarkup = tg.NewRemoveKeyboard(true)
		}
		chattableMessages = append(chattableMessages, message)
	}
	return chattableMessages
}

func (b *BaseBotContext) CreateAndSendMessage(message string) {
	messageConfigs := b.CreateMessages(message)
	var chattableMessages []tg.Chattable
	for _, msg := range messageConfigs {
		chattableMessages = append(chattableMessages, msg)
	}
	b.SendMessages(chattableMessages...)
}

func (b *BaseBotContext) SetKeyboard(keyboard *BotKeyboard) {
	b.DefaultKeyboard = keyboard
}

func (b *BaseBotContext) SendErrorMessage() {
	b.CreateAndSendMessage(b.ErrorMessage)
}

func (b *BaseBotContext) incCallCount() uint {
	b.CallCount++
	return b.CallCount
}
