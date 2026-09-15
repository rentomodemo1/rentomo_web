package pricing

import (
	"html/template"
	"net/http"
)

var screen = template.Must(template.New("pricing").Parse("<h1>Rental rates</h1><p>Daily: {{.Rates.Daily}}</p><p>Weekly: {{.Rates.Weekly}}</p><form method='get' action='/pricing'><label><input type='checkbox' name='annual' value='1' {{if .Annual}}checked{{end}}>Show annual rate</label><button>Update rates</button></form>{{if .Annual}}<p>Annual: {{.Rates.Annual}}</p>{{end}}"))

func Screen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	screen.Execute(w, struct {
		Rates  Rates
		Annual bool
	}{Lookup(), r.URL.Query().Get("annual") == "1"})
}
