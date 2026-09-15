package main

import (
	"github.com/rentomodemo1/rentomo_web/internal/booking"
	"log"
	"net/http"
	"time"
)

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/booking", booking.Summary)
	store := booking.NewStore(map[string]booking.Booking{"demo": {ID: "demo", FreeUntil: time.Now().Add(time.Hour)}}, func(string) error { return nil })
	mux.HandleFunc("/cancel", store.CancelHandler)

	return mux
}
func main() { log.Fatal(http.ListenAndServe(":8080", routes())) }
