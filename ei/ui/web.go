package ui

import (
	"fmt"
	app2 "github.com/imakiri/playground/pkg/app"
	"html/template"
	"net/http"
	"time"
)

type root struct {
	app *app2.App
}

// GET /
func (e root) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("/assets/css/style.css", nil)
		_ = p.Push("/assets/favicon.ico", nil)
	}
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		ise(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		ise(w, err)
		return
	}

	fmt.Printf("%v WebGetRoot endpoint hit by %s\n", time.Now(), r.RemoteAddr)
}
