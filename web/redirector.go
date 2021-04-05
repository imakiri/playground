package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/imakiri/gorum/web/internal"
	"log"
	"net"
	"net/http"
)

type HTTPSRedirector struct {
	server *http.Server
	status chan error
}

func (s *HTTPSRedirector) Launch() {
	var l, err = net.Listen("tcp", ":80")
	if err != nil {
		s.status <- err
	}

	go func() {
		s.status <- s.server.Serve(l)
	}()
}

func (s *HTTPSRedirector) Stop() {
	var err = s.server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func NewHTTPSRedirector(status chan error) (*HTTPSRedirector, error) {
	var r HTTPSRedirector
	var err error
	r.server = &http.Server{}
	r.status = status

	var router = mux.NewRouter()
	router.HandleFunc("/", internal.Go2HTTPS)
	r.server.Handler = router

	return &r, err
}
