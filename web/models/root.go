package models

import (
	"fmt"
	"html/template"
	"net/http"
)

type r0 struct{}

func (r0) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		_, _ = fmt.Fprintf(w, "Template error occurred: %s", err)
	}
	_ = t.ExecuteTemplate(w, "index", nil)
}

var Root r0
