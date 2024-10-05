package teleBotStateLib

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	var callCount = c.incCallCount()

	defer func() {
		if r := recover(); r != nil {
			if callCount == 1 {
				c.SendErrorMessage()
			}
			panic(r)
		}
	}()

	if callCount > MaxCallCount {
		panic(ToManyCalls)
	}

	currentState := m.StateManger.GetState(c.GetMessageSenderId())
	c.SetKeyboard(currentState.Keyboard)

	handlerResponse, isCommandProcessed = m.processCommand(c)
	if !isCommandProcessed {
		handlerResponse = m.defineNewState(c, currentState)
	}

	switch handlerResponse.TransitionType {
	case GoState:
		newState := handlerResponse.NextState
		err = m.transactToNewState(c, currentState, newState, false)
		if err != nil {
			panic(err)
		}
	case ReloadState:
		err = m.transactToNewState(c, currentState, currentState, false)
		if err != nil {
			panic(err)
		}
	case GoStateForce:
		newState := handlerResponse.NextState
		err = m.transactToNewState(c, currentState, newState, true)
		if err != nil {
			panic(err)
		}
	case GoStateInPlace:
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
	currentState *BotState,
) HandlerResponse {
	var handlerResponse HandlerResponse
	var buttonPressed bool

	if currentState.Keyboard != nil {
		handlerResponse, buttonPressed = currentState.Keyboard.ProcessMessage(c)
		if buttonPressed {
			return handlerResponse
		}
	}

	handlerResponse = currentState.Handler(c)
	return handlerResponse
}

func (m *BotStatesManager) transactToNewState(
	c BotContext,
	currentState *BotState,
	newState *BotState,
	forceTransition bool,
) error {
	var messages []string
	var err error

	if !forceTransition && currentState.MessageExit != nil {
		exitMessages, err := currentState.MessageExit.ToStringArray(c)
		if err != nil {
			panic(err)
		}
		messages = append(messages, exitMessages...)
	}

	if newState.MessageEnter != nil {
		enterMessages, err := newState.MessageEnter.ToStringArray(c)
		if err != nil {
			panic(err)
		}
		messages = append(messages, enterMessages...)
	}

	if len(messages) > 0 {
		var messagesForSend = c.CreateMessages(messages...)
		lastMessageForSend := &messagesForSend[len(messagesForSend)-1]
		if newState.Keyboard != nil {
			lastMessageForSend.ReplyMarkup = newState.Keyboard.GetKeyBoard()
		} else {
			lastMessageForSend.ReplyMarkup = tg.NewRemoveKeyboard(true)
		}

		var chattableMessages []tg.Chattable
		for _, msg := range messagesForSend {
			chattableMessages = append(chattableMessages, msg)
		}
		c.SendMessages(chattableMessages...)
	} else if newState.Keyboard != nil {
		log.Panicf("in state %s defined keyboard without enter message!", newState.BotStateName)
	}

	err = m.StateManger.SetState(c.GetMessageSenderId(), newState)
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
