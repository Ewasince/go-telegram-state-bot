package states

import (
	"github.com/Ewasince/go-telegram-state-bot/errors"

	//. "github.com/Ewasince/go-telegram-state-bot/types"
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
)

type baseBotState struct {
	BotStateName string
	MessageEnter Messagables
	MessageExit  Messagables
	Keyboard     Keyboard
	Handler      ContextHandler
}

func NewBotState(
	BotStateName string,
	MessageEnter Messagables,
	MessageExit Messagables,
	Keyboard Keyboard,
	Handler ContextHandler,
) BotState {
	if Keyboard != nil && MessageEnter == nil {
		panic(errors.KeyboardAndEnterMessage)
	}
	return baseBotState{
		BotStateName,
		MessageEnter,
		MessageExit,
		Keyboard,
		Handler,
	}
}

func (b baseBotState) GetBotStateName() string      { return b.BotStateName }
func (b baseBotState) GetMessageEnter() Messagables { return b.MessageEnter }
func (b baseBotState) GetMessageExit() Messagables  { return b.MessageExit }
func (b baseBotState) GetKeyboard() Keyboard        { return b.Keyboard }
func (b baseBotState) GetHandler() ContextHandler   { return b.Handler }
