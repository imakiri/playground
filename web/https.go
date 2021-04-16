package web

import (
	"context"
	"crypto/tls"
	"github.com/gorilla/mux"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/assets"
	"github.com/imakiri/gorum/types"
	"github.com/imakiri/gorum/utils"
	"golang.org/x/net/http2"
	"html/template"
	"log"
	"net"
	"net/http"
)

const (
	path_key  = "secrets/web/key.txt"
	path_cert = "secrets/web/certificate.crt"
)

type Services struct {
	Assets *assets.Service
}

type Service struct {
	https    bool
	server   *http.Server
	status   chan error
	index    *template.Template
	services Services
}

func register(s *Service) {
	var router = mux.NewRouter()
	router.Handle("/assets/", assets.Handler(s.services.Assets))
	router.HandleFunc("/", s.root)
	router.HandleFunc("/admin/load", s.load)
	s.server.Handler = router
}

func NewService(ss Services, status chan error, https bool) (*Service, error) {
	if utils.IsNil(ss) {
		return nil, erres.NilArgument.Extend(0).SetDescription("services cannot be nil")
	}

	var s Service
	var err error
	s.https = https
	s.server = &http.Server{}
	s.status = status
	s.services = ss

	if https {
		var cert tls.Certificate
		if cert, err = tls.LoadX509KeyPair(path_cert, path_key); err != nil {
			return nil, err
		}

		s.server.TLSConfig = &tls.Config{
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
	}

	if err = http2.ConfigureServer(s.server, nil); err != nil {
		return nil, err
	}

	register(&s)
	return &s, err
}

func (s *Service) Load() error {
	var assets, err = s.services.Assets.Get(context.Background(), &types.Request{})
	if err != nil {
		return err
	}

	s.assets = assets
	s.index, err = template.New("index").Parse(string(s.assets.Index))
	return err
}

func (s *Service) Launch() {
	var l net.Listener
	var err error

	if s.https {
		if l, err = tls.Listen("tcp", ":443", s.server.TLSConfig); err != nil {
			s.status <- err
		}
	} else {
		if l, err = net.Listen("tcp", ":80"); err != nil {
			s.status <- err
		}
	}

	go func() {
		s.status <- s.server.Serve(l)
	}()
}

func (s *Service) Stop() {
	var err error
	if err = s.server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
