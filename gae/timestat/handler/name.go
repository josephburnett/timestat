package handler

import (
	"fmt"
	"net/http"

	"timestat/datastore"

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
	timer, err := datastore.NewTimer(ctx, u.String(), name)
	if internalError(w, err) {
		return
	}
	fmt.Fprint(w, "New timer created with id: "+timer.ID)
}
