package teleBotStateLib

import (
	"github.com/Ewasince/go-telegram-state-bot/enums"
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
	"log"
)

const (
	MaxCallCount = 10
)

type BotStatesManager struct {
	BotCommands map[string]BotCommand
	StateManger StateCacheManager
}

func NewBotStatesManager(
	botCommands []BotCommand,
	stateManager StateCacheManager,
) *BotStatesManager {
	botCommandsMap := make(map[string]BotCommand, len(botCommands))
	for _, botCommand := range botCommands {
		botCommandsMap[botCommand.CommandMessage] = botCommand
	}

	return &BotStatesManager{
		BotCommands: botCommandsMap,
		StateManger: stateManager,
	}
}

func (m *BotStatesManager) ProcessMessage(c BotContext) {
	var err error
	var handlerResponse HandlerResponse
	var isCommandProcessed bool

	var callCount = c.IncCallCount()

	if callCount > MaxCallCount {
		panic(ToManyCalls)
	}

	currentState := *m.StateManger.GetState(c.GetMessageSenderId())
	c.SetKeyboard(currentState.GetKeyboard())

	handlerResponse, isCommandProcessed = m.processCommand(c)
	if !isCommandProcessed {
		handlerResponse = m.defineNewState(c, currentState)
	}

	switch handlerResponse.TransitionType {
	case enums.GoState:
		newState := handlerResponse.NextState
		err = m.transactToNewState(c, currentState, *newState, false)
		if err != nil {
			panic(err)
		}
	case enums.ReloadState:
		err = m.transactToNewState(c, currentState, currentState, false)
		if err != nil {
			panic(err)
		}
	case enums.GoStateForce:
		newState := handlerResponse.NextState
		err = m.transactToNewState(c, currentState, *newState, true)
		if err != nil {
			panic(err)
		}
	case enums.GoStateInPlace:
		err = m.StateManger.SetState(c.GetMessageSenderId(), handlerResponse.NextState)
		if err != nil {
			panic(err)
		}
		m.ProcessMessage(c)
	default:
	}
}

// defineNewState returns new bot state id, new state availability flag and error
func (m *BotStatesManager) defineNewState(
	c BotContext,
	currentState BotState,
) HandlerResponse {
	var handlerResponse HandlerResponse
	var buttonPressed bool

	if kb := currentState.GetKeyboard(); kb != nil {
		handlerResponse, buttonPressed = kb.ProcessMessage(c)
		if buttonPressed {
			return handlerResponse
		}
	}
	if handler := currentState.GetHandler(); handler != nil {
		handlerResponse = handler(c)
	} else {
		log.Print("No handler for " + currentState.GetBotStateName() + "!")
		handlerResponse = HandlerResponse{}
	}
	return handlerResponse
}

func (m *BotStatesManager) transactToNewState(
	c BotContext,
	currentState BotState,
	newState BotState,
	forceTransition bool,
) error {
	var messages []StateChattable
	var err error

	if extMsg := currentState.GetMessageExit(); !forceTransition && extMsg != nil {
		exitMessages, err := extMsg.ToTgMessages(c)
		if err != nil {
			panic(err)
		}
		messages = append(messages, exitMessages...)
	}

	if entMsg := newState.GetMessageEnter(); entMsg != nil {
		enterMessages, err := entMsg.ToTgMessages(c)
		if err != nil {
			panic(err)
		}
		messages = append(messages, enterMessages...)
	}

	if newKeyboard := newState.GetKeyboard(); len(messages) > 0 {
		c.SendChattables(newKeyboard, messages...)
	} else if newKeyboard != nil {
		log.Panicf("in state %s defined keyboard without enter message!", newState.GetBotStateName())
	}

	err = m.StateManger.SetState(c.GetMessageSenderId(), &newState)
	if err != nil {
		panic(err)
	}

	return nil
}

// processCommand returns new state, new state flag, command processed flag and err
func (m *BotStatesManager) processCommand(c BotContext) (HandlerResponse, bool) {
	botCommand, exists := m.BotCommands[c.GetMessageCommand()]
	if !exists {
		return HandlerResponse{}, false
	}
	handlerResponse := botCommand.CommandHandler(c)
	return handlerResponse, true
}
