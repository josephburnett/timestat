package handler

import "net/http"

func internalError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func notFound(w http.ResponseWriter) {
	http.Error(w, "404 Sorry there is nothing here.", http.StatusNotFound)
}
