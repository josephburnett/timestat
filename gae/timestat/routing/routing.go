package routing

import (
	"net/http"

	h "timestat/handler"
)

func init() {
	http.HandleFunc("/timer", h.Timer)
	http.HandleFunc("/start", h.Start)
	http.HandleFunc("/stop", h.Stop)
}
