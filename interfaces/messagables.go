package interfaces

type Messagables interface {
	ToTgMessages(c BotContext) ([]StateChattable, error)
}
