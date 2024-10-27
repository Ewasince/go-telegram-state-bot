package interfaces

type BotState interface {
	GetBotStateName() string
	GetMessageEnter() Messagables
	GetMessageExit() Messagables
	GetKeyboard() Keyboard
	GetHandler() ContextHandler
}
