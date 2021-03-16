package web

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func (s *Service) Root(w http.ResponseWriter, r *http.Request) {
	n := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] web.root hit by ip:%s\n", n, r.RemoteAddr)

	//if pusher, ok := w.(http.Pusher); ok {
	//	_ = pusher.Push("assets/css/style.css", nil)
	//	_ = pusher.Push("assets/favicon.ico", nil)
	//}

	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		s.ise(w, err)
		return
	}

	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		s.ise(w, err)
		return
	}
}
