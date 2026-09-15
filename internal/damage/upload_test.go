package damage

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"testing"
)

func TestCustomerUpload(t *testing.T) {
	var body bytes.Buffer
	form := multipart.NewWriter(&body)
	form.WriteField("booking", "b")
	photo, _ := form.CreateFormFile("photo", "damage.jpg")
	photo.Write([]byte("image bytes"))
	form.Close()
	r := httptest.NewRequest("POST", "/photos/upload", &body)
	r.Header.Set("Content-Type", form.FormDataContentType())
	w := httptest.NewRecorder()
	s := NewStore()
	s.UploadHandler(w, r)
	p, ok := s.Get("b")
	if w.Code != 201 || !ok || p.Name != "damage.jpg" || string(p.Data) != "image bytes" {
		t.Fatal(w.Code, p)
	}
}
func TestUploadRequiresPhotoAndBooking(t *testing.T) {
	s := NewStore()
	if s.Save(Photo{}) == nil {
		t.Fatal("empty photo accepted")
	}
	w := httptest.NewRecorder()
	s.UploadHandler(w, httptest.NewRequest("GET", "/photos/upload", nil))
	if w.Code != 405 {
		t.Fatal(w.Code)
	}
}
