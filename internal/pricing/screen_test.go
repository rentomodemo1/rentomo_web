package pricing

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRatesAndAnnualToggle(t *testing.T) {
	for _, annual := range []bool{false, true} {
		url := "/pricing"
		if annual {
			url += "?annual=1"
		}
		w := httptest.NewRecorder()
		Screen(w, httptest.NewRequest("GET", url, nil))
		body := w.Body.String()
		for _, want := range []string{"Daily: 50", "Weekly: 300", "method='get'", "name='annual'"} {
			if !strings.Contains(body, want) {
				t.Errorf("missing %s: %s", want, body)
			}
		}
		if strings.Contains(body, "Annual: 14000") != annual {
			t.Fatal("annual toggle does not change displayed rate")
		}
	}
}
