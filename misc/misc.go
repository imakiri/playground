package misc

import (
	"fmt"
	"net/http"
	"strings"
)

type Gyto struct {
	lik int
	lpe string
}

func (g *Gyto) Lik() *int {
	return &g.lik
}

func (g *Gyto) Lpe() *string {
	return &g.lpe
}

func Uuy(in string) (out string) {
	out = strings.ToLower(in)
	return
}

func testGo() {
	in := "ONJDFSGNJEGJNEGRF"

	// go re := misc.Uuy(in)

	ch := make(chan string)
	go func() {
		ch <- Uuy(in)
	}()
	re := <-ch

	//
	fmt.Print(re)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		_ = p.Push("style.css", nil)
	}

}
