package handler

import (
	"net/http"
	"timestat/datastore"

	"appengine"
	"appengine/user"
)

// Cancel deletes the current running timer.
func Cancel(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	timer, err := datastore.LoadRunningTimer(ctx, u.String())
	if internalError(w, err) {
		return
	}
	if timer == nil {
		printEmptyTimer(w)
		return
	}
	err = datastore.DeleteRunningTimer(ctx, timer)
	if internalError(w, err) {
		return
	}
	printEmptyTimer(w)
}
