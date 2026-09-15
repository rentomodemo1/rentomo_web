package damage

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http/httptest"
	"strings"
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
func TestAgentReviewStoredAgainstBooking(t *testing.T) {
	s := NewStore()
	s.Save(Photo{BookingID: "b", Name: "damage.jpg", Data: []byte("photo")})
	a := NewAttachments()
	reviews := NewReviews(s, a)
	r := httptest.NewRequest("POST", "/photos/review", strings.NewReader("booking=b&agent=alex&notes=scratch"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	reviews.Handler(w, r)
	if w.Code != 201 {
		t.Fatal(w.Code, w.Body.String())
	}
	got := httptest.NewRecorder()
	a.Handler(got, httptest.NewRequest("GET", "/booking/photos?booking=b", nil))
	var v Review
	if err := json.Unmarshal(got.Body.Bytes(), &v); err != nil || v.BookingID != "b" || v.Agent != "alex" || v.Notes != "scratch" {
		t.Fatal(v, err)
	}
}
