package examples

import (
	"github.com/Ewasince/go-telegram-state-bot"
	"github.com/Ewasince/go-telegram-state-bot/api_utils"
	"github.com/Ewasince/go-telegram-state-bot/helpers"
	"github.com/Ewasince/go-telegram-state-bot/interfaces"
	"github.com/Ewasince/go-telegram-state-bot/message_types"
	"github.com/Ewasince/go-telegram-state-bot/models"
	"github.com/Ewasince/go-telegram-state-bot/states"
	"log"
	"sync"

	tl "github.com/Ewasince/go-telegram-state-bot/context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var echoState = states.NewBotState(
	"Echo State",
	message_types.BotMessages{
		message_types.TextMessage("You entered to echo state!"),
		message_types.TextMessage("You might push buttons to reenter state"),
	},
	message_types.TextMessage("You exit from echo state"),
	nil,
	func(c interfaces.BotContext) interfaces.HandlerResponse {
		helpers.CreateAndSendMessage("You typed: "+c.GetMessageText(), c)
		return interfaces.HandlerResponse{}
	},
)

func MainBase() {
	bot, err := tgbotapi.NewBotAPI("YOUR_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	senderHandler := &api_utils.BaseSenderHandler{
		BotApi:   bot,
		BotMutex: &sync.Mutex{},
	}
	manager := teleBotStateLib.NewBotStatesManager(
		[]models.BotCommand{},
		teleBotStateLib.NewBaseStateCacheManager(&echoState),
	)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

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

		ctx := tl.NewBaseContext(update.Message, senderHandler)
		manager.ProcessMessage(ctx)
	}
}
