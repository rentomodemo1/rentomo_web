package booking

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSummary(t *testing.T) {
	w := httptest.NewRecorder()
	Summary(w, httptest.NewRequest("GET", "/booking", nil))
	if !strings.Contains(w.Body.String(), "Booking demo") {
		t.Fatal(w.Body.String())
	}
}
