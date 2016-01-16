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

func internalError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}
