package main

import (
	"fmt"
	"github.com/imakiri/playground/misc"
	"github.com/imakiri/playground/server"
	"net/http"
)

func main() {
	server.Run()
	//for i := 4; i > 0; i-- {
	//	app.RunTest1()
	//}
	//test()
}

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}

}

func test() {
	ch := make(chan string)
	in := "ONJDFSGNJEGJNEGRF"

	go func() {
		ch <- misc.Uuy(in)
	}()

	re := <-ch
	fmt.Print(re)

}
