package handler

import "net/http"

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

var menu, name, iden = "", "", ""

func init() {

	menu += "<h3>Menu</h3>"
	menu += "<ul>"
	menu += "<li><a href=\"/start\">start</a></li>"
	menu += "<li><a href=\"/timer\">timer</a></li>"
	menu += "<li><a href=\"/stop\">stop</a></li>"
	menu += "</ul>"

	name += "<h3>New Timer</h3>"
	name += "<form action=\"/name\" method=\"post\">"
	name += "<input type=\"text\" name=\"name\">"
	name += "<input type=\"submit\" value=\"Create\">"
	name += "</form>"

	iden += "<h3>Identify Timer</h3>"
	iden += "<form action=\"/identify\" method=\"post\">"
	iden += "<input type=\"text\" name=\"id\">"
	iden += "<input type=\"submit\" value=\"Identify\">"
	iden += "</form>"
}

func inHTMLBody(body string) string {
	return "<html><body>" + body + "</body></html>"
}

func messageHTML(message string) string {
	return "<p>" + message + "</p>"
}
