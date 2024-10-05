package teleBotStateLib

type BotMessages []string

func (b BotMessages) ToStringArray(c BotContext) ([]string, error) { return b, nil }

type BotMessageHandler func(c BotContext) ([]string, error)

func (b BotMessageHandler) ToStringArray(c BotContext) ([]string, error) { return b(c) }

type BotState struct {
	BotStateName string
	MessageEnter StringifyArray
	MessageExit  StringifyArray
	Keyboard     *BotKeyboard
	Handler      ContextHandler
}

func NewBotState(
	BotStateName string,
	MessageEnter StringifyArray,
	MessageExit StringifyArray,
	Keyboard *BotKeyboard,
	Handler ContextHandler,
) BotState {
	if Keyboard != nil && MessageEnter == nil {
		panic(KeyboardAndEnterMessage)
	}
	return BotState{
		BotStateName,
		MessageEnter,
		MessageExit,
		Keyboard,
		Handler,
	}
}
