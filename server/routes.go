package server

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api"
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
			Resolve(globalPlaces, api.V0, Sender, w, r)
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
					Resolve(globalPlaces, api.V1, Sender, w, r)
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
					Resolve(globalPlaces, api.V2, Sender, w, r)
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
					Resolve(globalPlaces, api.V3, Sender, w, r)
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
						Resolve(globalPlaces, api.V4, Sender, w, r)
					case "location":
						Resolve(globalPlaces, api.V4, Sender, w, r)
					case "visit":
						Resolve(globalPlaces, api.V4, Sender, w, r)
					}

				default:
					_, _ = io.WriteString(w, "Method is't implemented")
				}
			},
		},
	},
}
