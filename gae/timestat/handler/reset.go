package handler

import (
	"net/http"
	"timestat/datastore"
	m "timestat/model"
	"timestat/stats"

	"appengine"
)

// Reset collection running timer stats for a user and deletes the timer.
func Reset(w http.ResponseWriter, r *http.Request) {
	owner := r.FormValue("owner")
	if owner == "" {
		panic("Unknown owner.")
	}
	ctx := appengine.NewContext(r)
	timer, err := datastore.LoadRunningTimer(ctx, owner)
	if err != nil {
		panic("Unable to load running timer for " + owner + " : " + err.Error())
	}
	minute, err := stats.Minute(timer)
	if err != nil {
		panic(err)
	}
	err = updateAllDimensions(ctx, timer, minute)
	if err != nil {
		panic(err)
	}
	err = datastore.DeleteRunningTimer(ctx, timer)
	if err != nil {
		panic("Unable to delete running timer for " + owner + " : " + err.Error())
	}
}

func updateAllDimensions(ctx appengine.Context, timer *m.RunningTimer, minute int32) error {
	for _, dim := range m.AllDimensions {
		dist, err := loadOrCreateDistribution(ctx, timer, dim)
		if err != nil {
			return err
		}
		stats.Update(dist, minute)
		err = datastore.SaveDistribution(ctx, dist)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadOrCreateDistribution(ctx appengine.Context, timer *m.RunningTimer, dim m.Dimension) (*m.Distribution, error) {
	id, err := datastore.DimensionID(timer, dim)
	if err != nil {
		return nil, err
	}
	dist, err := datastore.LoadDistribution(ctx, timer.Owner, dim, id)
	if err != nil {
		return nil, err
	}
	if dist == nil {
		dist = datastore.NewDistribution(timer.Owner, dim, id, timer.TimerID)
	}
	return dist, nil
}
