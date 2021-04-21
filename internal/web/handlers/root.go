package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Root(templ *template.Template, w http.ResponseWriter, r *http.Request) error {
	var err error
	if pusher, ok := w.(http.Pusher); ok {
		if err = push(pusher); err != nil {
			fmt.Println(err)
		}
	}

	err = templ.ExecuteTemplate(w, "index", nil)
	if err != nil {
		return err
	}

	return err
}
