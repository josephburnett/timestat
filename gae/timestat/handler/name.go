package handler

import (
	"fmt"
	"net/http"

	"timestat/datastore"
	m "timestat/model"

	"appengine"
	"appengine/user"
)

// Name creates a new timer and associated it with the current running timer.
func Name(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name is required.", http.StatusBadRequest)
		return
	}
	running, err := datastore.LoadRunningTimer(ctx, u.String())
	if internalError(w, err) {
		return
	}
	if running == nil {
		userError(w, "There is no current timer.")
		return
	}
	if running.State == m.NamedRunning || running.State == m.NamedStopped {
		userError(w, "The current timer is already named.")
		return
	}
	timer, err := datastore.NewTimer(ctx, u.String(), name)
	if internalError(w, err) {
		return
	}
	running.TimerID = timer.ID
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
	fmt.Fprint(w, inHTMLBody(messageHTML("New timer created with id: "+timer.ID)+menu))
}
