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
	//http.Error(w, err, http.)
}
