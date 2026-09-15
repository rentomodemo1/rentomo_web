package main

import (
	"net/http/httptest"
	"testing"
)

func TestRegisteredRoutes(t *testing.T) {
	for _, p := range []string{"/booking"} {
		r := httptest.NewRecorder()
		routes().ServeHTTP(r, httptest.NewRequest("GET", p, nil))
		if r.Code == 404 {
			t.Fatalf("route missing: %s", p)
		}
	}
}
