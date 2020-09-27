package main

import (
	"github.com/imakiri/playground/server"
	"net/http"
)

func main() {
	server.Run()
	//for i := 4; i > 0; i-- {
	//	app.RunTest1()
	//}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}
}
