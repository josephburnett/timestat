package handler

import (
	"fmt"
	"net/http"
	"timestat/datastore"
	m "timestat/model"

	"appengine"
	"appengine/user"
)

// Identify associates the current running timer with a pre-existing timer.
func Identify(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	id := r.FormValue("id")
	if id == "" {
		userError(w, "ID is required.")
		return
	}
	running, err := datastore.LoadRunningTimer(ctx, u.String())
	if internalError(w, err) {
		return
	}
	if running == nil {
		userError(w, "No timer is running.")
		return
	}
	if running.State == m.NamedRunning || running.State == m.NamedStopped {
		userError(w, "Timer is already identified.")
		return
	}
	timer, err := datastore.LoadTimer(ctx, u.String(), id)
	if internalError(w, err) {
		return
	}
	if timer == nil {
		userError(w, "No such timer: "+id)
		return
	}
	running.TimerID = id
	if running.State == m.AnonRunning {
		running.State = m.NamedRunning
	}
	if running.State == m.AnonStopped {
		running.State = m.NamedStopped
	}
	err = datastore.SaveRunningTimer(ctx, running)
	if internalError(w, err) {
		return
	}
	fmt.Fprint(w, inHTMLBody(messageHTML("Time successfully identified as: "+timer.ID)+menu))
}
