package main

import (
	"github.com/imakiri/playground/db"
	"net/http"
)

func main() {
	db.Run()
	//server.Run()
	//goroutines.Vu()

}

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}
}
