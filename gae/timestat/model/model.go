package model

import (
	"time"
)

// Dimension is an enumerated set of dimensions along which Distributions can be
// kept.  E.g. "Day", "Month" or "DayOfTheWeek".
type Dimension string

// Valid Dimensions
const (
	General                      Dimension = "general"
	Day                                    = "day"
	Week                                   = "week"
	Month                                  = "month"
	Year                                   = "year"
	DayOfTheWeek                           = "day-of-the-week"
	TenMinuteTime                          = "ten-minute-time"
	TenMinuteTimeAndDayOfTheWeek           = "ten-minute-time-and-day-of-the-week"
)

// Distribution keeps stats for a given Dimension.  Distributions are
// partitioned by Owner.
type Distribution struct {
	Owner               string    // user name
	Dimension           Dimension //
	ID                  string    // e.g. "2015"
	TimerID             string    // foreign key
	MinuteProbabilities []float64 `datastore:",noindex"` // 1440 minutes
	PointCount          int       // number of timer instances in distribution
	Mean                float64   //
	Median              float64   //
	StandardDeviation   float64   //
}

// Timer is a thing that a user is timing.  E.g. "Daily commute" or "Fold the
// laundry".  Timers are partitioned by Owner.
type Timer struct {
	Owner    string    // user name
	ID       string    // stable url key
	Name     string    // user friendly name
	LastUsed time.Time //
}

// State for the singleton RunningTimer.
type State string

// Valid States.
const (
	RunningState State = "running" // initial state
	StoppedState       = "stopped" // ready to be reset
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

// Kind is a enumerated set of Datastore kinds.
type Kind string

// Valid Kinds.
const (
	DistributionKind Kind = "Distribution"
	TimerKind             = "Timer"
	RunningTimerKind      = "RunningTimer"
)
