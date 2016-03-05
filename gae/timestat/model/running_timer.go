package model

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	AnonRunning  State = "anon-running"  // initial state
	AnonStopped        = "anon-stopped"  // stopped but not associated with a timer id
	NamedRunning       = "named-running" // running and associated with a timer id
	NamedStopped       = "named-stopped" // ready to be reset
)

func (t *RunningTimer) Print(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	bytes, _ := json.Marshal(t)
	fmt.Fprint(w, string(bytes))
}
