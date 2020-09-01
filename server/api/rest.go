package api

import (
	"fmt"
	"github.com/gorilla/mux"
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
				case "POST":
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

func RunREST(rr *mux.Router) error {
	router := rr.PathPrefix("/api").Subrouter()

	for n, r := range Routs {
		router.HandleFunc(r.Path, r.Handler)

		fmt.Printf("Обработчик %s зарегестрирован на %s\n", n, r.Path)
	}

	return nil
}
