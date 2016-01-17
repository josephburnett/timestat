package handler

import (
	"fmt"
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
		fmt.Fprint(w, "No running timer.")
		return
	}
	fmt.Fprint(w, "Timer running since "+timer.Start.String())
}
