package misc

import (
	"fmt"
	"github.com/imakiri/playground/app"
	"github.com/imakiri/playground/data/inside"
	"github.com/imakiri/playground/data/schema"
	"log"
	"net/http"
	"reflect"
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

func Test6() {
	var c = inside.GetUserV1{Data: &schema.User{Login: "imari"}}

	if err := inside.Exec(c).ExecuteSQL(); err != nil {
		//fmt.Print(err.Error())
		fmt.Print(reflect.TypeOf(err))
	} else {
		fmt.Print(c.Data.Name)
	}
}

func Test7() {
	var c = inside.DeleteUserV1{Data: &schema.User{Login: "imri"}}

	if err := inside.Exec(c).ExecuteSQL(); err != nil {
		//fmt.Print(err.Error())
		fmt.Print(reflect.TypeOf(err))
	} else {
		fmt.Print(c.Data)
	}
}

func Test8() {
	var c = inside.GetUserV2{Data: &schema.User{Login: "imakiri", Id: 7}}

	if err := inside.Exec(c).ExecuteSQL(); err != nil {
		fmt.Print(err.Error() + "\n")
		fmt.Print(reflect.TypeOf(err))
	} else {
		fmt.Print(c.Data)
	}
}
