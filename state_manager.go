package teleBotStateLib

import (
	. "github.com/Ewasince/go-telegram-state-bot/interfaces"
)

type baseStateCacheManager struct {
	StatesCache  map[int64]*BotState
	DefaultState *BotState
}

func NewBaseStateCacheManager(defaultState *BotState) StateCacheManager {
	return &baseStateCacheManager{
		StatesCache:  map[int64]*BotState{},
		DefaultState: defaultState,
	}
}

func (s *baseStateCacheManager) SetState(key int64, BaseBotState *BotState) error {
	s.StatesCache[key] = BaseBotState
	return nil
}

func (s *baseStateCacheManager) GetState(key int64) *BotState {
	state, exists := s.StatesCache[key]
	if !exists {
		return s.DefaultState
	}
	return state
}
