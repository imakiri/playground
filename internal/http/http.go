package http

import (
	"context"
	"github.com/gorilla/mux"
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
	router.HandleFunc("/", go2https)
	r.server.Handler = router

	return &r, err
}

func go2https(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}
