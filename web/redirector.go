package web

import (
	"context"
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
	r.server = &http.Server{}
	r.status = status

	var err = internal.RedirectorHTTPS(r.server)
	if err != nil {
		return nil, err
	}

	return &r, err
}
