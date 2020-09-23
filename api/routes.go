package api

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/api/endpoints"
	"github.com/imakiri/playground/app/remote"
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
	Handler:    nil,
	Routs:      []Route{},
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
					Resolve(endpoints.Location, []remote.Caster{remote.PlacePhotos}, w, r)
				case "location":
					Resolve(endpoints.Location, []remote.Caster{remote.PlacePhotos}, w, r)
				case "visit":
					Resolve(endpoints.Location, []remote.Caster{remote.PlacePhotos}, w, r)
				}
			},
		},
	},
}
