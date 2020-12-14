package api

import (
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/app"
	"github.com/imakiri/playground/core"
)

// Returns prepared router and redirectionRouter for http api server, or error
func NewAPIRouters(s core.Settings) (*mux.Router, *mux.Router, error) {
	_, err := app.NewApp(s)
	if err != nil {
		return nil, nil, err
	}

	var router = mux.NewRouter()
	var redirRouter = mux.NewRouter()

	//redirRouter.HandleFunc("/", redirect)
	//
	//router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	//router.Handle("/", root{a})
	//router.Handle("/detect", detect{a})

	return router, redirRouter, nil
}
