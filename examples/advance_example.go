package examples

import (
	"fmt"
	"github.com/Ewasince/go-telegram-state-bot"
	. "github.com/Ewasince/go-telegram-state-bot/api_utils"
	. "github.com/Ewasince/go-telegram-state-bot/context"
	. "github.com/Ewasince/go-telegram-state-bot/enums"
	"github.com/Ewasince/go-telegram-state-bot/helpers"
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
	. "github.com/Ewasince/go-telegram-state-bot/keyboard"
	. "github.com/Ewasince/go-telegram-state-bot/message_types"
	. "github.com/Ewasince/go-telegram-state-bot/models"
	"github.com/Ewasince/go-telegram-state-bot/states"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"runtime/debug"
	"sync"
)

var defaultState = &echoAdvancedState
var defaultKeyboard = &echoKeyboard

var startCommand = BotCommand{
	CommandMessage: "start",
	CommandHandler: func(c BotContext) HandlerResponse {
		helpers.CreateAndSendMessage("Hello!", c)

		return HandlerResponse{
			NextState:      defaultState,
			TransitionType: GoStateForce,
		}
	},
}

var echoKeyboard = BotKeyboard{Keyboard: []ButtonsRow{
	{
		BotButton{
			ButtonTitle:   "Left Top corner",
			ButtonHandler: keyboardEchoHandler,
		},
		BotButton{
			ButtonTitle:   "Right Top corner",
			ButtonHandler: keyboardEchoHandler,
		},
	},
	{
		BotButton{
			ButtonTitle:   "Left Bottom corner",
			ButtonHandler: keyboardEchoHandler,
		},
		BotButton{
			ButtonTitle:   "Right Bottom corner",
			ButtonHandler: keyboardEchoHandler,
		},
	},
}}

func keyboardEchoHandler(_ BotContext) HandlerResponse {
	return HandlerResponse{
		TransitionType: ReloadState,
	}
}

var echoAdvancedState = states.NewBotState(
	"Echo State",
	BotMessages{
		TextMessage("You entered to echo state!"),
		TextMessage("You might push buttons to reenter state"),
	},
	TextMessage("You exit from echo state"),
	defaultKeyboard,
	func(c BotContext) HandlerResponse {
		helpers.CreateAndSendMessage("You typed: "+c.GetMessageText(), c)
		return HandlerResponse{}
	},
)

func getProcessFunc(sender *BaseSenderHandler) func(*tg.Message) {
	cache := teleBotStateLib.NewBaseStateCacheManager(defaultState)
	manager := teleBotStateLib.NewBotStatesManager(
		[]BotCommand{
			startCommand,
		},
		cache,
	)

	return func(message *tg.Message) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error occurred when handle message: " + r.(error).Error() + "\n" + string(debug.Stack()))
			}
		}()

		ctx := NewBaseContext(message, sender)
		manager.ProcessMessage(ctx)
	}
}

func MainAdvanced() {
	botAPI, err := tg.NewBotAPI("YOUR_TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	senderHandler := &BaseSenderHandler{
		BotApi:   botAPI,
		BotMutex: &sync.Mutex{},
	}

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := botAPI.GetUpdatesChan(u)

	processMessage := getProcessFunc(senderHandler)
	fmt.Println("Bot started...")
	for update := range updates {
		if update.Message == nil {
			continue
		}

		messageMessage := update.Message
		messageSender := messageMessage.From

		log.Printf(
			"[%s, %d] %s",
			messageSender.UserName,
			messageSender.ID,
			update.Message.Text,
		)

		processMessage(update.Message)
	}

}
