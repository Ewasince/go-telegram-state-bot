package teleBotStateLib

type StateCacheManager interface {
	SetState(int64, *BotState) error
	GetState(int64) *BotState
}

type BaseStateCacheManager struct {
	StatesCache  map[int64]*BotState
	DefaultState *BotState
}

func NewBaseStateCacheManager(defaultState *BotState) StateCacheManager {
	return &BaseStateCacheManager{
		StatesCache:  map[int64]*BotState{},
		DefaultState: defaultState,
	}
}

func (s *BaseStateCacheManager) SetState(key int64, botState *BotState) error {
	s.StatesCache[key] = botState
	return nil
}

func (s *BaseStateCacheManager) GetState(key int64) *BotState {
	state, exists := s.StatesCache[key]
	if !exists {
		return s.DefaultState
	}
	return state
}
