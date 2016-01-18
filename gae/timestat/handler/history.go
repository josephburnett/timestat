package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"timestat/datastore"
	m "timestat/model"

	"appengine"
	"appengine/user"
)

// History handles anything under "/history/".  The rest of the URL specifies
// the timer id, dimension and id.  E.g. "/history/daily-commute/year/2015"
func History(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	path := r.URL.Path[len("/history/"):]
	elements := strings.Split(path, "/")
	if len(elements) != 3 {
		http.NotFound(w, r)
		return
	}
	timerID, dimension, id := elements[0], m.Dimension(elements[1]), elements[2]
	dist, err := datastore.LoadDistribution(ctx, u.String(), timerID, dimension, id)
	if internalError(w, err) {
		return
	}
	if dist == nil {
		http.NotFound(w, r)
		return
	}
	bytes, err := json.Marshal(dist)
	if internalError(w, err) {
		return
	}
	fmt.Fprint(w, inHTMLBody(messageHTML("History: "+path)+string(bytes)+menu))
}
