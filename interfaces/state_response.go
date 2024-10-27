package interfaces

import (
	. "github.com/Ewasince/go-telegram-state-bot/enums"
)

type HandlerResponse struct {
	NextState      *BotState // which state should go next
	TransitionType StateTransitionType
}

// ContextHandler returns new state id, is new state flag and error
type ContextHandler func(c BotContext) HandlerResponse
