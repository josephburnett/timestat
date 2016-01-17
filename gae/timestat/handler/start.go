package handler

import (
	"fmt"
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
		fmt.Fprint(w, "There is already a timer running.")
		return
	}
	timer, err = datastore.NewRunningTimer(ctx, u.String())
	if internalError(w, err) {
		return
	}
	fmt.Fprint(w, "Successfully started a timer.")
}
