package datastore

import (
	"net/url"
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
		key := datastore.NewKey(c, m.RunningTimerKind, timer.Owner, 0, nil)
		err := datastore.Delete(c, key)
		return err
	}, nil)
	return err
}
