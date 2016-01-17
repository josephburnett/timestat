package handler

import (
	"net/http"
	"timestat/datastore"

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
		panic("Unable to load running timer for " + owner)
	}
	// TODO collect stats on timer
	err = datastore.DeleteRunningTimer(ctx, timer)
	if err != nil {
		panic("Unable to delete running timer for " + owner)
	}
}
