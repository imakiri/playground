package web

import (
	"fmt"
	"net/http"
	"time"
)

func (s *webService) root(w http.ResponseWriter, r *http.Request) {
	var t = time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] / hit by ip:%s\n", t, r.RemoteAddr)

	w.Header().Set("Content-Type", "text/html")

	if pusher, ok := w.(http.Pusher); ok {
		var err = push(pusher)
		if err != nil {
			fmt.Println(err)
		}
	}

	var err = s.template.home.ExecuteTemplate(w, "root", nil)
	if err != nil {
		ise(w, err)
		return
	}
}
