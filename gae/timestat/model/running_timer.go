package model

import (
	"time"
)

// RunningTimer is a singleton entity representing a user collecting timing
// data.  The existence
type RunningTimer struct {
	Owner   string    // user name
	State   State     //
	Start   time.Time //
	End     time.Time //
	TimerID string    // foreign key
}

// State for the singleton RunningTimer.
type State string

// Valid States.
const (
	RunningState State = "running" // initial state
	StoppedState       = "stopped" // ready to be reset
)
