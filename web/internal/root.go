package internal

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func Root(w http.ResponseWriter, r *http.Request) error {
	var err error
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] web.root hit by ip:%s\n", t, r.RemoteAddr)

	if pusher, ok := w.(http.Pusher); ok {
		if err = pusher.Push("/assets/css", nil); err != nil {
			fmt.Println(err)
		}
		if err = pusher.Push("/assets/ico", nil); err != nil {
			fmt.Println(err)
		}
	}

	w.Header().Set("Content-Type", "text/html")

	var templ *template.Template
	templ, err = template.ParseFiles("web/templates/index.html")
	if err != nil {
		return err
	}

	err = templ.ExecuteTemplate(w, "index", nil)
	if err != nil {
		return err
	}

	return err
}
