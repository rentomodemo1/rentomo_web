package main

import (
	"github.com/rentomodemo1/rentomo_web/internal/booking"
	"log"
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/booking", booking.Summary)

	return mux
}
func main() { log.Fatal(http.ListenAndServe(":8080", routes())) }
