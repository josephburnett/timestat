package handler

import (
	"net/http"
	"timestat/datastore"

	"appengine"
)

func internalError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func userError(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

// TODO: create a templating component for rendering page views.

var menu, name = "", ""

func init() {

	menu += "<h3>Menu</h3>"
	menu += "<ul>"
	menu += "<li><a href=\"/start\">start</a></li>"
	menu += "<li><a href=\"/stop\">stop</a></li>"
	menu += "<li><a href=\"/timer\">timer</a></li>"
	menu += "<li><a href=\"/cancel\">cancel</a></li>"
	menu += "</ul>"

	name += "<h3>New Timer</h3>"
	name += "<form action=\"/name\" method=\"post\">"
	name += "<input type=\"text\" name=\"name\">"
	name += "<input type=\"submit\" value=\"Create\">"
	name += "</form>"
}

func inHTMLBody(body string) string {
	return "<html><body>" + body + "</body></html>"
}

func messageHTML(message string) string {
	return "<p>" + message + "</p>"
}

func timersHTML(c appengine.Context, owner string) (string, error) {
	timers, err := datastore.ListTimers(c, owner)
	if err != nil {
		return "", err
	}
	list := "<h3>Identify timer</h3>"
	list += "<ol>"
	for _, t := range timers {
		list += "<li><a href=\"/identify?id=" + t.ID + "\">" + t.Name + "</a></li>"
	}
	list += "</ol>"
	return list, nil
}
