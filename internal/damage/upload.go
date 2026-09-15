package damage

import (
	"errors"
	"io"
	"net/http"
	"sync"
)

type Photo struct {
	BookingID, Name string
	Data            []byte
}
type Store struct {
	mu     sync.Mutex
	photos map[string]Photo
}

func NewStore() *Store { return &Store{photos: map[string]Photo{}} }
func (s *Store) Get(id string) (Photo, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.photos[id]
	p.Data = append([]byte(nil), p.Data...)
	return p, ok
}
func (s *Store) Save(p Photo) error {
	if p.BookingID == "" || len(p.Data) == 0 {
		return errors.New("booking and photo required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	p.Data = append([]byte(nil), p.Data...)
	s.photos[p.BookingID] = p
	return nil
}
func (s *Store) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "use POST", 405)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 2<<20)
	if err := r.ParseMultipartForm(2 << 20); err != nil {
		http.Error(w, "invalid photo upload", 400)
		return
	}
	defer r.MultipartForm.RemoveAll()
	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "photo required", 400)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "cannot read photo", 400)
		return
	}
	if err = s.Save(Photo{r.FormValue("booking"), header.Filename, data}); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(201)
	w.Write([]byte("Photo uploaded"))
}
