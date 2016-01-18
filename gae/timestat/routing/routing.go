package routing

import (
	"net/http"

	h "timestat/handler"
)

func init() {
	http.HandleFunc("/timer", h.Timer)
	http.HandleFunc("/start", h.Start)
	http.HandleFunc("/stop", h.Stop)
	http.HandleFunc("/name", h.Name)
	http.HandleFunc("/history/", h.History)
	http.HandleFunc("/task/reset", h.Reset)
}
