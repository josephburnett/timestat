package datastore

import (
	"regexp"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"

	m "timestat/model"
)

// NewTimer creates a new Time entity.
func NewTimer(owner, name string) *m.Timer {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	id := reg.ReplaceAllString(name, "-")
	id = strings.ToLower(id)
	timer := &m.Timer{
		Owner:    owner,
		ID:       id,
		Name:     name,
		LastUsed: time.Now(),
	}
	return timer
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
