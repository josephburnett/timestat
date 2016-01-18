package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"timestat/datastore"
	m "timestat/model"

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
	bytes, _ := json.Marshal(timer)
	fmt.Fprint(w, inHTMLBody(messageHTML("Running timer:")+string(bytes)+menu))
	if timer.State == m.AnonRunning || timer.State == m.AnonStopped {
		fmt.Fprint(w, name+iden)
	}
}
