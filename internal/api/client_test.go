package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJSON(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("{\"count\":2}")) }))
	defer s.Close()
	var got struct{ Count int }
	if err := GetJSON(s.Client(), s.URL, &got); err != nil || got.Count != 2 {
		t.Fatalf("%+v %v", got, err)
	}
}
