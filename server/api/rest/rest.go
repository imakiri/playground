package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api"
	"github.com/imakiri/playground/server/storage"
	"io"
	"net/http"
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
					w1 := make(chan api.Thing)
					//w2 := make(chan Thing)
					go api.Api.GetThing(&storage.Local, "example", w1)
					//go api.Api.GetThing(, "example", w2)

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
		router.HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик view зарегестрирован на %s\n", r.Path)
	}

	for _, r := range Action.Routs {
		router.HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик action зарегестрирован на %s\n", r.Path)
	}

	return nil
}
