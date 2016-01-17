package handler

import (
	"fmt"
	"net/http"
	"timestat/datastore"

	"appengine"
	"appengine/user"
)

// Stop terminates a running timer if one exists.
func Stop(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	timer, err := datastore.LoadRunningTimer(ctx, u.String())
	if internalError(w, err) {
		return
	}
	if timer == nil {
		fmt.Fprint(w, "There is not a timer running.")
		return
	}
	err = datastore.StopRunningTimer(ctx, timer)
	if internalError(w, err) {
		return
	}
	fmt.Fprint(w, "Timer successfully stopped.")
}
