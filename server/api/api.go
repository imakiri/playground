package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/app"
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

type Places map[string]interfaces.Api

var App app.App

var globalPlaces = Places{
	"local": &app.Local,
}

func Resolve(p Places, resolver interfaces.Resolver, w http.ResponseWriter, r *http.Request) {
	wg := sync.WaitGroup{}
	l := len(p)

	c := make(chan interfaces.Thing, l)
	defer close(c)
	wg.Add(l)

	for k, p := range p {
		go func(k string, p interfaces.Api) {
			defer wg.Done()
			interfaces.Api.GetThing(p, k, c)
		}(k, p)
	}
	wg.Wait()

	resolver(interfaces.Parcel{
		Channel:        c,
		Request:        r,
		ResponseWriter: w,
	})
}

func RunREST(rr *mux.Router) error {
	router := rr.PathPrefix("/api").Subrouter()

	router.HandleFunc(View.PrefixPath, View.Handler)
	router.HandleFunc(Action.PrefixPath, Action.Handler)

	for _, r := range View.Routs {
		router.PathPrefix("/view").Subrouter().HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик view зарегестрирован на %s\n", r.Path)
	}

	for _, r := range Action.Routs {
		router.PathPrefix("/action").Subrouter().HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик action зарегестрирован на %s\n", r.Path)
	}

	return nil
}
