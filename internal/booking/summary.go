package booking

import (
	"html/template"
	"net/http"
)

var summary = template.Must(template.New("booking").Parse("<h1>Booking {{.}}</h1><form method='post' action='/cancel'><input type='hidden' name='booking' value='{{.}}'><button>Cancel booking</button></form>"))

func Summary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	summary.Execute(w, "demo")
}
