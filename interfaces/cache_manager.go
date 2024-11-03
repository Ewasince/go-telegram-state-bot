package interfaces

type StateCacheManager interface {
	SetState(int64, *BotState) error
	GetState(int64) *BotState
}
