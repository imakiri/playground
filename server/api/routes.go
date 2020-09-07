package api

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api/resolvers"
	"github.com/imakiri/playground/server/core"
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
		Resolve(globalPlaces, resolvers.Location, core.SenderImp, w, r)
	},
	Routs: []Route{
		{
			Path: "/user/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				Resolve(globalPlaces, resolvers.Location, core.SenderImp, w, r)
			},
		},
		{
			Path: "/location",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				Resolve(globalPlaces, resolvers.Location, core.SenderImp, w, r)
			},
		},
		{
			Path: "/location/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				Resolve(globalPlaces, resolvers.Location, core.SenderImp, w, r)
			},
		},
		{
			Path: "/visit/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				Resolve(globalPlaces, resolvers.Location, core.SenderImp, w, r)
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
				switch mux.Vars(r)["entity"] {
				case "user":
					Resolve(globalPlaces, resolvers.V4, core.SenderImp, w, r)
				case "location":
					Resolve(globalPlaces, resolvers.V4, core.SenderImp, w, r)
				case "visit":
					Resolve(globalPlaces, resolvers.V4, core.SenderImp, w, r)
				}
			},
		},
	},
}
