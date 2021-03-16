package api

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/service"
)

type Service struct {
	service.Service
	config *cfg.Api
}

func New(bs service.Service) (*Service, error) {
	var s Service
	var err error

	s.Service = bs
	s.config, err = s.Cfg().Get4Api(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

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
