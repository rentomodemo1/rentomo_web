package damage

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Review struct{ BookingID, Agent, Notes string }
type Attachments struct {
	mu      sync.Mutex
	reviews map[string]Review
}

func NewAttachments() *Attachments   { return &Attachments{reviews: map[string]Review{}} }
func (a *Attachments) Save(r Review) { a.mu.Lock(); defer a.mu.Unlock(); a.reviews[r.BookingID] = r }
func (a *Attachments) Get(id string) (Review, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	r, ok := a.reviews[id]
	return r, ok
}
func (a *Attachments) Handler(w http.ResponseWriter, r *http.Request) {
	v, ok := a.Get(r.URL.Query().Get("booking"))
	if !ok {
		http.Error(w, "booking review not found", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
