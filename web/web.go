package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/transport"
	"net/http"
)

type Service struct {
	//gate        gate.GeneralService
	cc          transport.CfgClient
	config      *core.CfgWeb
	Server      *http.Server
	RedirServer *http.Server
}

func NewService(cc transport.CfgClient) error {
	var s Service
	var err error

	s.cc = cc
	s.config, err = s.cc.Web(context.Background(), &core.Request{})
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	redirRouter := mux.NewRouter()

	redirRouter.HandleFunc("/", redirect)

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", s.Root)

	s.Server = &http.Server{}
	s.RedirServer = &http.Server{}
	s.Server.Handler = router
	s.RedirServer.Handler = redirRouter

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

// Redirect from HTTP to HTTPS
func redirect(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}

// Internal ServiceInt Error Response
func ise(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(err.Error()))
}
