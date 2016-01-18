package datastore

import (
	"appengine"
	"appengine/datastore"

	m "timestat/model"
)

// NewDistribution creates a new distribution in the owner's namespace with the
// given dimension and id.  The new distribution is intialized with 1440 minute
// probabilities and zero point count.
func NewDistribution(owner string, dimension m.Dimension, id string, timerID string) *m.Distribution {
	return &m.Distribution{
		Owner:        owner,
		Dimension:    dimension,
		ID:           id,
		TimerID:      timerID,
		MinuteCounts: make([]int32, 1440),
		SampleCount:  0,
	}
}

// LoadDistribution loads a distribution from Datastore if it exists
func LoadDistribution(c appengine.Context, owner string, dimension m.Dimension, id string) (*m.Distribution, error) {
	key := datastore.NewKey(c, string(Distribution), compositeDistributionKey(owner, string(dimension), id), 0, nil)
	dist := &m.Distribution{}
	err := datastore.Get(c, key, dist)
	if err == datastore.ErrNoSuchEntity {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return dist, nil
}

// SaveDistribution saves a Distribution to Datastore.
func SaveDistribution(c appengine.Context, dist *m.Distribution) error {
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		distKey := compositeDistributionKey(dist.Owner, string(dist.Dimension), dist.ID)
		key := datastore.NewKey(c, string(Distribution), distKey, 0, nil)
		_, err := datastore.Put(c, key, dist)
		return err
	}, nil)
	return err
}

func compositeDistributionKey(owner, dimension, id string) string {
	return owner + "$" + dimension + "$" + id
}
