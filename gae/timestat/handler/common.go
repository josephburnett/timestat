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

var menu = ""

func init() {

	menu += "<ul>"
	menu += "<li><a href=\"/start\">start</a></li>"
	menu += "<li><a href=\"/timer\">timer</a></li>"
	menu += "<li><a href=\"/stop\">stop</a></li>"
	menu += "</ul>"
}

func inHTMLBody(body string) string {
	return "<html><body>" + body + "</body></html>"
}

func messageHTML(message string) string {
	return "<p>" + message + "</p>"
}
