package handler

import (
	"fmt"
	"net/http"
)

// Root is the homepage for the app.
func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, inHTMLBody(messageHTML("Time something!")+menu))
}
