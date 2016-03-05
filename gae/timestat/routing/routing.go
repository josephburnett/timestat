package routing

import (
	"net/http"

	h "timestat/handler"
)

func init() {
	http.HandleFunc("/", h.Root)
	http.HandleFunc("/timer", h.Timer)
	http.HandleFunc("/timer/start", h.Start)
	http.HandleFunc("/timer/stop", h.Stop)
	http.HandleFunc("/timer/name", h.Name)
	http.HandleFunc("/timer/identify", h.Identify)
	http.HandleFunc("/timer/cancel", h.Cancel)
	http.HandleFunc("/history/", h.History)
	http.HandleFunc("/task/reset", h.Reset)
}
