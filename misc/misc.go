package misc

import (
	"fmt"
	"github.com/imakiri/playground/app"
	"github.com/imakiri/playground/data"
	"github.com/imakiri/playground/data/inside"
	"github.com/imakiri/playground/data/schema"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func Test2() {
	switch err := data.Init().(type) {
	case data.InitError:
		log.Fatal(err.Error())
	}

	i := data.Internal()
	var u = &schema.User{Login: "imri"}

	switch err := i.GetUser(u).(type) {
	case inside.NotFoundError:
		fmt.Print("Not found")
	case inside.IncorrectArgumentError:
		fmt.Print("Incorrect argument")
	case error:
		fmt.Printf("Error: %v", reflect.TypeOf(err))
	default:

	}

	fmt.Printf("%s\n", u.Name)
}

func Test3() {
	switch err := data.Init().(type) {
	case data.InitError:
		log.Fatal(err.Error())
	}

	i := data.Internal()
	var u = &schema.User{Login: "imari", Name: "klff", Avatar: []byte("erhg"), PassHash: []byte("efge")}

	switch err := i.CreateUser(u).(type) {
	case inside.NotFoundError:
		fmt.Print("Not found")
	case inside.IncorrectArgumentError:
		fmt.Print("Incorrect argument")
	case error:
		fmt.Printf("Error: %v", err.Error())
	default:

	}
}

func Test4() {
	_, err := ioutil.ReadFile("data/dsn")

	switch err.(type) {
	case *os.PathError:
		fmt.Print("false")
	default:
		fmt.Print("default")
	}
}

func Test5() {
	e := func() (err error) {
		return nil
	}

	switch e().(type) {
	case NotFound:
		print("ffh\n")
	default:
		print("hjk\n")
	}
}

func Test6() {

}

var Er = E{
	NotFound:        NotFound{},
	ServiceInternal: ServiceInternal{},
}

type E struct {
	NotFound        NotFound
	ServiceInternal ServiceInternal
}

type NotFound struct {
}

func (NotFound) Error() string {
	return ""
}

type ServiceInternal struct {
}

func (ServiceInternal) Error() string {
	return ""
}
