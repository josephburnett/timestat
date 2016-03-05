package handler

import (
	"net/http"
	"timestat/datastore"
	m "timestat/model"

	"appengine"
	"appengine/user"
)

// Stop terminates a running timer if one exists.  If the running timer has not
// yet been associated with a timer id, it is transitioned to the AnonStopped
// state.  Otherwise it is transitioned to the NamedStopped state.
func Stop(w http.ResponseWriter, r *http.Request) {
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
	if timer.State == m.AnonStopped || timer.State == m.NamedStopped {
		timer.Print(w)
		return
	}
	if timer.TimerID == "" {
		timer.State = m.AnonStopped
	} else {
		timer.State = m.NamedStopped
	}
	err = datastore.SaveRunningTimer(ctx, timer)
	if internalError(w, err) {
		return
	}
	timer.Print(w)
}
