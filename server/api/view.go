package api

import (
	"io"
	"net/http"
)

var View = RootRoute{
	PrefixPath: "/view",
	Handler: func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			Resolve(globalPlaces, App.Do0, w, r)
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
					Resolve(globalPlaces, App.Do1, w, r)
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
					Resolve(globalPlaces, App.Do2, w, r)
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
					Resolve(globalPlaces, App.Do3, w, r)
				default:
					_, _ = io.WriteString(w, "Method is't implemented")
				}
			},
		},
	},
}
