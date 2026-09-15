package booking

import (
	"errors"
	"net/http"
	"sync"
	"time"
)

type Booking struct {
	ID        string
	FreeUntil time.Time
	Cancelled bool
}
type Store struct {
	mu       sync.Mutex
	bookings map[string]Booking
	refund   func(string) error
	Now      func() time.Time
}

func NewStore(b map[string]Booking, refund func(string) error) *Store {
	return &Store{bookings: b, refund: refund, Now: time.Now}
}
func (s *Store) Cancel(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.bookings[id]
	if !ok {
		return errors.New("booking not found")
	}
	if b.Cancelled {
		return nil
	}
	if s.Now().After(b.FreeUntil) {
		return errors.New("free cancellation window has ended")
	}
	if err := s.refund(id); err != nil {
		return err
	}
	b.Cancelled = true
	s.bookings[id] = b
	return nil
}
func (s *Store) CancelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "use POST", 405)
		return
	}
	if err := s.Cancel(r.FormValue("booking")); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte("Booking cancelled and refund raised"))
}
