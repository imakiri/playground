package web

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	path_key  = "secrets/web/key.txt"
	path_cert = "secrets/web/certificate.crt"
)

type Service struct {
	Server      *http.Server
	RedirServer *http.Server
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

func NewService() (*Service, error) {
	var cert, err = tls.LoadX509KeyPair(path_cert, path_key)
	if err != nil {
		return nil, err
	}

	var s Service
	s.Server = &http.Server{}
	s.RedirServer = &http.Server{}
	s.Server.TLSConfig = &tls.Config{
		Rand:                  nil,
		Time:                  nil,
		Certificates:          []tls.Certificate{cert},
		GetCertificate:        nil,
		GetClientCertificate:  nil,
		GetConfigForClient:    nil,
		VerifyPeerCertificate: nil,
		VerifyConnection:      nil,
		RootCAs:               nil,
		NextProtos:            nil,
		ServerName:            "",
		ClientAuth:            0,
		ClientCAs:             nil,
		InsecureSkipVerify:    false,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
		},
		PreferServerCipherSuites:    false,
		SessionTicketsDisabled:      false,
		ClientSessionCache:          nil,
		MinVersion:                  tls.VersionTLS12,
		MaxVersion:                  0,
		CurvePreferences:            nil,
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}

	err = registerRouts(&s)
	if err != nil {
		return nil, err
	}

	return &s, err
}

func (s *Service) Launch() error {
	var rsc = make(chan error)
	var sc = make(chan error)

	go func(rsc chan error) {
		rsc <- s.RedirServer.ListenAndServe()
	}(rsc)

	var l, err = tls.Listen("tcp", ":443", s.Server.TLSConfig)
	if err != nil {
		return err
	}

	go func(sc chan error) {
		sc <- s.Server.Serve(l)
	}(sc)

	select {
	case err = <-rsc:
		_ = s.Server.Shutdown(context.Background())
		return err
	case err = <-sc:
		_ = s.RedirServer.Shutdown(context.Background())
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
