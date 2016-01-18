package datastore

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"

	m "timestat/model"
)

// Linear probe for a timer id based on name.
func getTimerID(c appengine.Context, owner, name string) (string, error) {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	id := reg.ReplaceAllString(name, "-")
	id = strings.ToLower(id)
	for i := 1; i < 10; i++ {
		if i > 1 {
			id = id[:len(id)-1] + string(i)
		}
		timer, err := LoadTimer(c, owner, id)
		if err != nil {
			return "", err
		}
		if timer == nil {
			return id, nil
		}
	}
	return "", errors.New("There are already too many of those. Name it something else.")
}

// NewTimer creates a new Timer entity.  New timers are associated with the
// current running timer.
func NewTimer(c appengine.Context, owner, name string) (*m.Timer, error) {
	timer := &m.Timer{
		Owner:    owner,
		Name:     name,
		LastUsed: time.Now(),
	}
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		running, err := LoadRunningTimer(c, owner)
		if running.State != m.AnonRunning && running.State != m.AnonStopped {
			return errors.New("There is no currently running timer.")
		}
		id, err := getTimerID(c, owner, name)
		if err != nil {
			return err
		}
		timer.ID = id
		timerKey := compositeTimerKey(timer.Owner, timer.ID)
		key := datastore.NewKey(c, Timer, timerKey, 0, nil)
		_, err = datastore.Put(c, key, timer)
		if err != nil {
			return err
		}
		running.TimerID = id
		// TODO: State transitions should not be in the data layer.
		//       But I need to transactionally create a new timer,
		//       update the running timer and transition its state.
		if running.State == m.AnonRunning {
			running.State = m.NamedRunning
		}
		if running.State == m.AnonStopped {
			running.State = m.NamedStopped
			err = resetTimer(c, owner)
			if err != nil {
				return err
			}
		}
		err = SaveRunningTimer(c, running)
		return err
	}, nil)
	if err != nil {
		return nil, err
	}
	return timer, nil
}

// LoadTimer loads an existing Timer from Datastore.
func LoadTimer(c appengine.Context, owner, id string) (*m.Timer, error) {
	key := datastore.NewKey(c, Timer, compositeTimerKey(owner, id), 0, nil)
	timer := &m.Timer{}
	err := datastore.Get(c, key, timer)
	if err == datastore.ErrNoSuchEntity {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return timer, nil
}

// SaveTimer saves a Timer to Datastore.
func SaveTimer(c appengine.Context, timer *m.Timer) error {
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		timerKey := compositeTimerKey(timer.Owner, timer.ID)
		key := datastore.NewKey(c, Timer, timerKey, 0, nil)
		_, err := datastore.Put(c, key, timer)
		return err
	}, nil)
	return err
}

func compositeTimerKey(owner, id string) string {
	return owner + "$" + id
}
