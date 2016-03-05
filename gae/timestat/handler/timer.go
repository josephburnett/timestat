package handler

import (
	"net/http"
	"timestat/datastore"

	"appengine"
	"appengine/user"
)

// Timer renders the current running timer for a user.
func Timer(w http.ResponseWriter, r *http.Request) {
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
	timer.Print(w)
}
