package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/service"
	"net/http"
)

type Service struct {
	service.Service
	config      *cfg.Web
	Server      *http.Server
	RedirServer *http.Server
}

func New(bs service.Service) (*Service, error) {
	var s Service
	var err error

	s.Service = bs
	s.config, err = s.Cfg().Get4Web(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()
	redirRouter := mux.NewRouter()

	redirRouter.HandleFunc("/", s.redirect)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", s.Root)

	s.Server = &http.Server{}
	s.RedirServer = &http.Server{}
	s.Server.Handler = router
	s.RedirServer.Handler = redirRouter

	return &s, err
}

func (s *Service) Launch() error {
	var err error

	rsc := make(chan error)
	sc := make(chan error)

	go func(rsc chan error) {
		rsc <- s.RedirServer.ListenAndServe()
	}(rsc)

	go func(sc chan error) {
		sc <- s.Server.ListenAndServeTLS(s.config.CertFile, s.config.KeyFile)
	}(sc)

	select {
	case err = <-rsc:
		return err
	case err = <-sc:
		return err
	}
}

func (Service) redirect(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}

func (Service) ise(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
