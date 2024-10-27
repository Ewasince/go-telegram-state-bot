package interfaces

type BotContext interface {
	GetMessageCommand() string
	GetMessageText() string
	GetMessageSenderId() int64
	GetMessageChatId() int64

	SendChattables(Keyboard, ...StateChattable)

	//CreateMessages(...string) []*tg.MessageConfig
	//CreatePhotos(...ImgMessagePRIVATEEEEEEEEEEEEEEE) []*tg.PhotoConfig

	//CreateAndSendMessage(string)

	SetKeyboard(Keyboard)

	//SendErrorMessage()
	IncCallCount() uint
}
