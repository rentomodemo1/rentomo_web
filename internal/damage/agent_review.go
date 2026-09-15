package damage

import (
	"net/http"
)

type Reviews struct {
	photos      *Store
	attachments *Attachments
}

func NewReviews(p *Store, a *Attachments) *Reviews { return &Reviews{p, a} }
func (s *Reviews) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "use POST", 405)
		return
	}
	id := r.FormValue("booking")
	agent := r.FormValue("agent")
	notes := r.FormValue("notes")
	if _, ok := s.photos.Get(id); !ok || agent == "" || notes == "" {
		http.Error(w, "photo, agent and notes required", 400)
		return
	}
	s.attachments.Save(Review{id, agent, notes})
	w.WriteHeader(201)
	w.Write([]byte("Review saved against booking"))
}
