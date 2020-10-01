package misc

import (
	"fmt"
	"github.com/imakiri/playground/app"
	"github.com/imakiri/playground/data"
	"github.com/imakiri/playground/data/inside"
	"github.com/imakiri/playground/data/schema"
	"log"
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

func Test1() {
	err := app.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 4; i > 0; i-- {
		app.RunTest1()
	}
}

func Test2() {
	switch data.Init().(type) {
	case data.InitError:

	}

	i := data.Internal()
	var u = &schema.User{Login: "imari"}

	switch i.GetUser(u).(type) {
	case inside.NotFoundError:
		fmt.Print("Not found")
	case inside.IncorrectArgumentError:
		fmt.Print("Incorrect argument")
	default:

	}

	if err := i.GetUser(u); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%s\n", u.Name)
}
