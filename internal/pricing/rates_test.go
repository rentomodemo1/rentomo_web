package pricing

import "testing"

func TestRates(t *testing.T) {
	r := Lookup()
	if r.Daily <= 0 || r.Weekly <= 0 {
		t.Fatal(r)
	}
}
