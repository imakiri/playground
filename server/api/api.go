package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api/endpoints"
	"github.com/imakiri/playground/server/app/methods/remote"
	"github.com/imakiri/playground/server/core"
	"net/http"
	"sync"
)

func Resolve(resolver endpoints.Resolver, casters []remote.Caster, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	l := len(casters)

	wg := sync.WaitGroup{}
	wg.Add(l)
	c := make(chan core.ThingImp, l)
	defer close(c)

	for _, caster := range casters {
		go caster.Cast(&wg, c)
	}
	wg.Wait()

	resolver.Resolve(core.Parcel{
		Channel:        &c,
		Request:        r,
		ResponseWriter: w,
	})
}

func Run(rr *mux.Router) error {
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
