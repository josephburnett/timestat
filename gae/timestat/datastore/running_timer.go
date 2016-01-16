package datastore

import (
	"time"

	"appengine"
	"appengine/datastore"

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
		key := datastore.NewKey(c, m.RunningTimerKind, timer.Owner, 0, nil)
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
	key := datastore.NewKey(c, m.RunningTimerKind, owner, 0, nil)
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
		key := datastore.NewKey(c, m.RunningTimerKind, timer.Owner, 0, nil)
		// TODO: persist reset task
		// _, err := datastore.Put(c, key, timer)
		err := datastore.Delete(c, key)
		return err
	}, nil)
	return err
}
