package teleBotStateLib

type StateTransitionType uint

const (
	// DontGoState is default value. dont do any transition
	DontGoState StateTransitionType = iota
	// GoState set new state and wait new update
	GoState StateTransitionType = iota
	// GoStateInPlace set new state and immediately pass control to his handler
	GoStateInPlace StateTransitionType = iota //
	// GoStateForce set new state without exit message
	GoStateForce StateTransitionType = iota //
	// ReloadState acts like where is transition from another state to current
	ReloadState StateTransitionType = iota
)
