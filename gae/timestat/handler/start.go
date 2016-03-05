package handler

import (
	"net/http"
	"timestat/datastore"

	"appengine"
	"appengine/user"
)

// Start creates a running timer if one doesn't exist.
func Start(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	timer, err := datastore.LoadRunningTimer(ctx, u.String())
	if internalError(w, err) {
		return
	}
	if timer != nil {
		timer.Print(w)
		return
	}
	timer, err = datastore.NewRunningTimer(ctx, u.String())
	if internalError(w, err) {
		return
	}
	timer.Print(w)
}
