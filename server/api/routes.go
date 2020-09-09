package api

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/server/api/resolvers"
	"github.com/imakiri/playground/server/data/remote/casters"

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
		Resolve(resolvers.Location, []casters.Caster{casters.PlacePhotos}, w, r)
	},
	Routs: []Route{
		{
			Path: "/user/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				Resolve(resolvers.Location, []casters.Caster{casters.PlacePhotos}, w, r)
			},
		},
		{
			Path: "/location",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				Resolve(resolvers.Location, []casters.Caster{casters.PlacePhotos}, w, r)
			},
		},
		{
			Path: "/location/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				Resolve(resolvers.Location, []casters.Caster{casters.PlacePhotos}, w, r)
			},
		},
		{
			Path: "/visit/{id}",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				Resolve(resolvers.Location, []casters.Caster{casters.PlacePhotos}, w, r)
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
					Resolve(resolvers.Location, []casters.Caster{casters.PlacePhotos}, w, r)
				case "location":
					Resolve(resolvers.Location, []casters.Caster{casters.PlacePhotos}, w, r)
				case "visit":
					Resolve(resolvers.Location, []casters.Caster{casters.PlacePhotos}, w, r)
				}
			},
		},
	},
}
