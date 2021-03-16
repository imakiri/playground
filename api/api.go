package api

import (
	"github.com/imakiri/gorum/cfg"
)

type Service struct {
	cfg    cfg.ServiceClient
	config cfg.Config
}

func NewService(cfgc cfg.ServiceClient) (*Service, error) {
	var s Service
	var err error

	//API = s.API
	//_, err := app.NewApp(s)
	//if err != nil {
	//	return nil, nil, err
	//}
	//
	//var router = mux.NewRouter()
	//var redirRouter = mux.NewRouter()
	//
	////redirRouter.HandleFunc("/", redirect)
	////
	////router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))
	////router.Handle("/", root{a})
	////router.Handle("/detect", detect{a})

	return &s, err
}
