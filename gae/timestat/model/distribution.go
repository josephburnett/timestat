package model

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
