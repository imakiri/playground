package models

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type r0 struct{}

func (r0) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("assets/css/style.css", nil)
	}

	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("web/templates/index.html")

	if err != nil {
		_, _ = fmt.Fprintf(w, "Template error occurred: %s", err)
	}
	_ = t.ExecuteTemplate(w, "index", nil)

	fmt.Printf("%v Web/Root passed to %s\n", time.Now(), r.RemoteAddr)
}

var Root r0
