package model

import (
	"time"
)

type Dimension string

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

type Distribution struct {
	Owner               string    // user name
	Dimension           Dimension //
	Id                  string    // e.g. "2015"
	TimerId             string    // foreign key
	MinuteProbabilities []float64 `datastore:",noindex"` // 1440 minutes
	PointCount          int       // number of timer instances in distribution
	Mean                float64   //
	Median              float64   //
	StandardDeviation   float64   //
}

type Timer struct {
	Owner    string    // user name
	Id       string    // stable url key
	Name     string    // user friendly name
	LastUsed time.Time //
}

type State string

const (
	Reset   State = "reset"   // ready to start
	Running       = "running" //
	Stopped       = "stopped" // ready to be reset
)

type RunningTimer struct {
	Owner   string    // user name
	State   State     //
	Start   time.Time //
	End     time.Time //
	TimerId           // foreign key
}
