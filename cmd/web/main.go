package main

import (
	"github.com/rentomodemo1/rentomo_web/internal/booking"
	"github.com/rentomodemo1/rentomo_web/internal/damage"
	"log"
	"net/http"
	"time"
)

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/booking", booking.Summary)
	store := booking.NewStore(map[string]booking.Booking{"demo": {ID: "demo", FreeUntil: time.Now().Add(time.Hour)}}, func(string) error { return nil })
	mux.HandleFunc("/cancel", store.CancelHandler)

	photos := damage.NewStore()
	mux.HandleFunc("/photos/upload", photos.UploadHandler)
	attachments := damage.NewAttachments()
	reviews := damage.NewReviews(photos, attachments)
	mux.HandleFunc("/photos/review", reviews.Handler)
	mux.HandleFunc("/booking/photos", attachments.Handler)
	return mux
}
func main() { log.Fatal(http.ListenAndServe(":8080", routes())) }
