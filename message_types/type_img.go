package message_types

import (
	"errors"
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
	"github.com/Ewasince/go-telegram-state-bot/state_chattable"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

var _ Messagables = (*imgMessage)(nil) // interface hint

type imgMessage struct {
	ImagePath string
	ImageName string
}

func NewImgMessage(imagePath, imageName string) Messagables {
	if _, err := os.Stat("/path/to/whatever"); errors.Is(err, os.ErrNotExist) {
		panic("Not found image for path: " + imagePath)
	}
	return &imgMessage{
		ImagePath: imagePath,
		ImageName: imageName,
	}
}

func (m imgMessage) ToTgMessages(c BotContext) ([]StateChattable, error) {
	photoBytes, err := os.ReadFile(m.ImagePath)
	if err != nil {
		return nil, err
	}

	message := tg.NewPhoto(c.GetMessageChatId(), tg.FileBytes{
		Name:  m.ImageName,
		Bytes: photoBytes,
	})
	stateChattable := state_chattable.NewBaseStateChattable(&message, &message.BaseChat)

	return []StateChattable{stateChattable}, nil
}
