package booking

import (
	"errors"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestCancelAndRefund(t *testing.T) {
	now := time.Unix(1000, 0)
	calls := 0
	s := NewStore(map[string]Booking{"b": {ID: "b", FreeUntil: now.Add(time.Hour)}}, func(id string) error {
		if id != "b" {
			t.Fatal(id)
		}
		calls++
		return nil
	})
	s.Now = func() time.Time { return now }
	for range 2 {
		r := httptest.NewRequest("POST", "/cancel", strings.NewReader(url.Values{"booking": {"b"}}.Encode()))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()
		s.CancelHandler(w, r)
		if w.Code != 200 {
			t.Fatal(w.Code, w.Body.String())
		}
	}
	if calls != 1 || !s.bookings["b"].Cancelled {
		t.Fatal("cancellation or refund missing", calls)
	}
}
func TestRejectExpiredAndRefundFailure(t *testing.T) {
	now := time.Unix(1000, 0)
	calls := 0
	s := NewStore(map[string]Booking{"old": {FreeUntil: now.Add(-time.Second)}, "new": {FreeUntil: now.Add(time.Hour)}}, func(string) error { calls++; return errors.New("refund unavailable") })
	s.Now = func() time.Time { return now }
	if s.Cancel("old") == nil || calls != 0 {
		t.Fatal("expired cancellation accepted")
	}
	if s.Cancel("new") == nil || s.bookings["new"].Cancelled {
		t.Fatal("refund failure ignored")
	}
}
