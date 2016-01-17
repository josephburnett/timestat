package model

import (
	"time"
)

// Timer is a thing that a user is timing.  E.g. "Daily commute" or "Fold the
// laundry".  Timers are partitioned by Owner.
type Timer struct {
	Owner    string    // user name
	ID       string    // stable url key
	Name     string    // user friendly name
	LastUsed time.Time //
}
