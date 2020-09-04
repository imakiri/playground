package server

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

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
						Resolve(globalPlaces, App.Do4, w, r)
					case "location":
						Resolve(globalPlaces, App.Do4, w, r)
					case "visit":
						Resolve(globalPlaces, App.Do4, w, r)
					}

				default:
					_, _ = io.WriteString(w, "Method is't implemented")
				}
			},
		},
	},
}
