package datastore

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/taskqueue"

	m "timestat/model"
)

// NewRunningTimer start and persists a running time now.
func NewRunningTimer(c appengine.Context, owner string) (*m.RunningTimer, error) {
	timer := &m.RunningTimer{
		Owner: owner,
		State: m.RunningState,
		Start: time.Now(),
	}
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		key := datastore.NewKey(c, RunningTimer, timer.Owner, 0, nil)
		_, err := datastore.Put(c, key, timer)
		return err
	}, nil)
	if err != nil {
		return nil, err
	}
	return timer, nil
}

// LoadRunningTimer loads the current running timer for a user if one exists.
func LoadRunningTimer(c appengine.Context, owner string) (*m.RunningTimer, error) {
	key := datastore.NewKey(c, RunningTimer, owner, 0, nil)
	timer := &m.RunningTimer{}
	err := datastore.Get(c, key, timer)
	if err == datastore.ErrNoSuchEntity {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return timer, nil
}

// StopRunningTimer sets the running timer into a stopped state and transaction-
// ally inserts a task to collect statistics and delete the timer.
func StopRunningTimer(c appengine.Context, timer *m.RunningTimer) error {
	timer.State = m.StoppedState
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		timer.End = time.Now()
		key := datastore.NewKey(c, RunningTimer, timer.Owner, 0, nil)
		_, err := datastore.Put(c, key, timer)
		if err != nil {
			return err
		}
		t := taskqueue.NewPOSTTask("/task/reset", url.Values{
			"owner": []string{timer.Owner},
		})
		_, err = taskqueue.Add(c, t, "")
		return err
	}, nil)
	return err
}

// DeleteRunningTimer deletes a running timer unconditionally.
func DeleteRunningTimer(c appengine.Context, timer *m.RunningTimer) error {
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		key := datastore.NewKey(c, RunningTimer, timer.Owner, 0, nil)
		err := datastore.Delete(c, key)
		return err
	}, nil)
	return err
}

// DimensionID extracts the id for the running timer for the given
// dimension.  E.g. "monday" or "2015"
func DimensionID(timer *m.RunningTimer, dim m.Dimension) (string, error) {
	start := timer.Start
	if start.IsZero() {
		return "", errors.New("Timer is not started.")
	}
	format := fmt.Sprintf
	lower := strings.ToLower
	tenMinuteTime := func() string {
		hour := start.Hour()
		minute := start.Minute()
		minute = minute / 10 * 10
		return format("%02d-%02d", hour, minute)
	}
	dayOfTheWeek := func() string {
		return lower(start.Weekday().String())
	}
	switch {
	case dim == m.General:
		return "general", nil
	case dim == m.Day:
		return start.Format("2006-01-02"), nil
	case dim == m.Week:
		year, week := start.ISOWeek()
		return format("%v-W%02d", year, week), nil
	case dim == m.Month:
		return lower(start.Month().String()), nil
	case dim == m.Year:
		return format("%v", start.Year()), nil
	case dim == m.DayOfTheWeek:
		return dayOfTheWeek(), nil
	case dim == m.TenMinuteTime:
		return tenMinuteTime(), nil
	case dim == m.TenMinuteTimeAndDayOfTheWeek:
		return format("%v-%v", dayOfTheWeek(), tenMinuteTime()), nil
	}
	return "", errors.New("Unknown dimension.")
}
