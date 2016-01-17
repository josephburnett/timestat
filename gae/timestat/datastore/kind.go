package datastore

// Kind is a enumerated set of Datastore kinds.
type Kind string

// Valid Kinds.
const (
	Distribution Kind = "Distribution"
	Timer             = "Timer"
	RunningTimer      = "RunningTimer"
)
