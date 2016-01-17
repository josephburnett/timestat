package datstore

import (
	"appengine"
	"appengine/datastore"

	m "timestat/model"
)

// NewDistribution creates a new distribution in the owner's namespace with the
// given dimension and id.  The new distribution is intialized with 1440 minute
// probabilities and zero point count.
func NewDistribution(owner string, dimension Dimension, id string, timerID string) *m.Distribution {
	dist := &Distribution{
		Owner:               owner,
		Dimension:           dimension,
		ID:                  id,
		TimerId:             timerID,
		MinuteProbabilities: make([]float64, 1440),
		PointCount:          0,
	}
	for i := range dist.MinuteProbabilities {
		dist.MinuteProbabilities[i] = 1.0
	}
	// TODO: initialize the derived properties of Distribution and normalize
	return dist
}

// LoadDistribution loads a distribution from Datastore if it exists
func LoadDistribution(c appengine.Context, owner string, dimension Dimension, id string) {
	key := datastore.NewKey(c, Distribution, compositeDistributionKey(owner, dimension, id), 0, nil)
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
		distKey := compositeDistributionKey(owner, dimension, id)
		key := datastore.NewKey(c, Distribution, distKey, 0, nil)
		_, err := datastore.Put(c, key, dist)
		return err
	}, nil)
	return err
}

func compositeDistributionKey(owner, dimension, id string) string {
	return owner + "$" + dimension + "$" + id
}
