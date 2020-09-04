package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api"
	"github.com/imakiri/playground/server/app/apiApp"
	"github.com/imakiri/playground/server/storage"
	"io"
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

var wg = sync.WaitGroup{}

var App apiApp.App

var View = RootRoute{
	PrefixPath: "/view",
	Handler: func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			q := r.URL.Query()
			_, _ = io.WriteString(w,
				fmt.Sprintf("User: %v, Location: %v, Visit: %v",
					q.Get("user"),
					q.Get("location"),
					q.Get("visit")))

		default:
			_, _ = io.WriteString(w, "Method is't implemented")
		}
	},
	Routs: []Route{
		{
			Path: "/user/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
					places := []api.Api{
						&storage.Local,
					}
					l := len(places)

					c := make(chan api.Thing, l)
					defer close(c)
					wg.Add(l)

					for _, p := range places {
						go func(p api.Api) {
							defer wg.Done()
							api.Api.GetThing(p, "example", c)
						}(p)
					}
					wg.Wait()

					App.MatchUp(api.Parcel{
						Channel:        c,
						Request:        r,
						ResponseWriter: w,
					})

					_, _ = io.WriteString(w, "Done")

				default:
					_, _ = io.WriteString(w, "Method is't implemented")
				}
			},
		},
		{
			Path: "/location/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
				case "POST":
				default:
					_, _ = io.WriteString(w, "Method is't implemented")
				}
			},
		},
		{
			Path: "/visit/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
				case "POST":
				default:
					_, _ = io.WriteString(w, "Method is't implemented")
				}
			},
		},
	},
}

var Action = RootRoute{
	PrefixPath: "/action",
	Handler:    nil,
	Routs: []Route{
		{
			Path: "/add/{entity}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "POST":
					switch mux.Vars(r)["entity"] {
					case "user":
					case "location":
					case "visit":
					}

				default:
					_, _ = io.WriteString(w, "Method is't implemented")
				}
			},
		},
	},
}

func RunREST(rr *mux.Router) error {
	router := rr.PathPrefix("/api").Subrouter()

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
