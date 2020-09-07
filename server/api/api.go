package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/core"
	"net/http"
	"sync"
)

var globalPlaces = core.Places{}

func Resolve(
	places core.Places,
	resolver core.Resolver,
	sender core.Sender,
	w http.ResponseWriter,
	r *http.Request,
) {

	wg := sync.WaitGroup{}
	l := len(places)

	c := make(chan core.Thing, l)

	defer close(c)
	wg.Add(l)

	for k, p := range places {
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

func RegisterApiHandlers(rr *mux.Router) error {
	router := rr.PathPrefix("/api").Subrouter()

	viewRouter := router.Methods("GET").Subrouter()
	actionRouter := router.Methods("POST").Subrouter()

	viewRouter.HandleFunc(View.PrefixPath, View.Handler)
	fmt.Printf("Обработчик /view зарегистрирован\n")

	actionRouter.HandleFunc(Action.PrefixPath, Action.Handler)
	fmt.Printf("Обработчик /action зарегистрирован\n")

	for _, r := range View.Routs {
		viewRouter.PathPrefix("/view").Subrouter().HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик view зарегистрирован на %s\n", r.Path)
	}

	for _, r := range Action.Routs {
		actionRouter.PathPrefix("/action").Subrouter().HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик action зарегистрирован на %s\n", r.Path)
	}
	return nil
}
