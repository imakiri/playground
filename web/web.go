package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
	"net/http"
)

type Config interface {
	Get4Web(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.Web, error)
}

type Service struct {
	config       Config
	configCached *cfg.Web
	Server       *http.Server
	RedirServer  *http.Server
}

func registerRouts(s *Service) error {
	var forum *HandlerForum
	var err error

	forum, err = newHandlerForum()
	if err != nil {
		return err
	}

	var router = mux.NewRouter()
	var redirRouter = mux.NewRouter()

	redirRouter.HandleFunc("/", s.redirect)
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", s.Root)
	router.Handle("/forum/", forum)

	s.Server.Handler = router
	s.RedirServer.Handler = redirRouter

	return err
}

func NewService(c Config) (*Service, error) {
	var s Service
	var err error

	s.config = c
	s.configCached, err = s.config.Get4Web(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	s.Server = &http.Server{}
	s.RedirServer = &http.Server{}

	err = registerRouts(&s)
	if err != nil {
		return nil, err
	}

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
		sc <- s.Server.ListenAndServeTLS(s.configCached.CertFile, s.configCached.KeyFile)
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
