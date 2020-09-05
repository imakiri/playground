package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api"
	"github.com/imakiri/playground/server/interfaces"
	"net/http"
	"sync"
)

type Route struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
}

type RootRoute struct {
	PrefixPath string
	Handler    func(w http.ResponseWriter, r *http.Request)
	Routs      []Route
}

var globalPlaces = interfaces.Places{
	"local": &api.Local,
}

func Resolve(p interfaces.Places, resolver interfaces.Resolver, sender interfaces.Sender, w http.ResponseWriter, r *http.Request) {
	wg := sync.WaitGroup{}
	l := len(p)

	c := make(chan interfaces.Thing, l)
	defer close(c)
	wg.Add(l)

	for k, p := range p {
		go func(k string, p interfaces.Api) {
			defer wg.Done()
			sender.Send(p, k, c)
		}(k, p)
	}
	wg.Wait()

	resolver.Resolve(interfaces.Parcel{
		Channel:        c,
		Request:        r,
		ResponseWriter: w,
	})
}

func RegisterApiHandlers(rr *mux.Router) error {
	router := rr.PathPrefix("/api").Subrouter()

	router.HandleFunc(View.PrefixPath, View.Handler)
	fmt.Printf("Обработчик /view зарегистрирован")

	router.HandleFunc(Action.PrefixPath, Action.Handler)
	fmt.Printf("Обработчик /action зарегистрирован")

	for _, r := range View.Routs {
		router.PathPrefix("/view").Subrouter().HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик view зарегистрирован на %s\n", r.Path)
	}

	for _, r := range Action.Routs {
		router.PathPrefix("/action").Subrouter().HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик action зарегистрирован на %s\n", r.Path)
	}
	return nil
}
