# Go Telegram State Bot


This package was developed for building telegram bots based on state logic

## Getting started

### Prerequisites

**Go Telegram State Bot** requires [Go](https://go.dev/) version [1.22](https://go.dev/doc/devel/release#go1.22.0) or above.

### Getting Go Telegram State Bot

With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), `go [build|run|test]` automatically fetches the necessary dependencies when you add the import in your code:

```sh
import "github.com/Ewasince/go-telegram-state-bot"
```

Alternatively, use `go get`:

```sh
go get github.com/Ewasince/go-telegram-state-bot
```

### Running Go Telegram State Bot

A basic example:

```go
package main


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

func main() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
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

```

This example also located in [Base Example](examples/base_example.go)

To run the code, use the `go run` command, like:

```sh
$ go run example.go
```

Type `/start` or any other text to your bot for example

### See more examples

#### Examples

Advanced example located in [Advanced Example](examples/advance_example.go), 
which includes besides state also keyboard, command and more other features. 

## Documentation

Documentation will appear when at least one person start using this module. 
Currently, all functionality can be understood based on [Advanced Example](examples/advance_example.go).