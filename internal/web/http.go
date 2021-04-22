package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/imakiri/gorum/internal/web/handlers"
	"log"
	"net"
	"net/http"
)

type Redirector struct {
	server *http.Server
	status chan error
}

func (s *Redirector) Launch() {
	var l, err = net.Listen("tcp", ":80")
	if err != nil {
		s.status <- err
	}

	go func() {
		s.status <- s.server.Serve(l)
	}()
}

func (s *Redirector) Stop() {
	var err = s.server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func NewRedirector(status chan error) (*Redirector, error) {
	var r Redirector
	var err error
	r.server = &http.Server{}
	r.status = status

	var router = mux.NewRouter()
	router.HandleFunc("/", handlers.Go2HTTPS)
	r.server.Handler = router

	return &r, err
}
