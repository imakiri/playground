package goroutines

import (
	"github.com/imakiri/playground/goroutines/a"
	"github.com/imakiri/playground/goroutines/inter"
	"github.com/imakiri/playground/server/core"
	"net/http"
	"sync"
)

var N a.Foo

var D a.Bar

func Vu() {
	inter.Bar.Does(D, N)
}

func Resolve(p core.Places, resolver core.Resolver, sender core.Sender, w http.ResponseWriter, r *http.Request) {
	wg := sync.WaitGroup{}
	l := len(p)

	//////////////////////////////////////////////////////////////
	// Сомнительная конструкция
	c := make(chan core.Thing, l)
	//////////////////////////////////////////////////////////////
	defer close(c)
	wg.Add(l)

	for k, p := range p {
		go func(k string, p core.Api) {
			defer wg.Done()
			sender.Send(p, k, c)
		}(k, p)
	}
	wg.Wait()

	resolver.Resolve(core.Parcel{
		Channel:        c,
		Request:        r,
		ResponseWriter: w,
	})
}

var globalPlaces = core.Places{}
