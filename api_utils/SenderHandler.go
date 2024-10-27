package api_utils

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
)

type SenderHandler interface {
	SendChattable(tg.Chattable) error
}

type BaseSenderHandler struct {
	BotApi   *tg.BotAPI
	BotMutex *sync.Mutex
}

func (b *BaseSenderHandler) SendChattable(msg tg.Chattable) error {
	b.BotMutex.Lock()
	defer b.BotMutex.Unlock()
	if _, err := b.BotApi.Send(msg); err != nil {
		log.Panic(err)
		return err
	}
	return nil
}
