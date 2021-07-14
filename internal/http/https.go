package http

import (
	"context"
	"crypto/tls"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/utils"
	"golang.org/x/net/http2"
	"log"
	"net"
	"net/http"
)

const (
	path_key  = "secrets/web/key.txt"
	path_cert = "secrets/web/certificate.crt"
)

type Server struct {
	https  bool
	server *http.Server
	status chan error
}

func NewServer(web http.Handler, status chan error, https bool) (*Server, error) {
	if utils.IsNil(web) {
		return nil, erres.NilArgument.Extend(0).SetDescription("web cannot be nil")
	}

	var s = new(Server)
	s.server = new(http.Server)
	s.server.Handler = web
	s.https = https
	s.status = status

	if https {
		var cert, err = tls.LoadX509KeyPair(path_cert, path_key)
		if err != nil {
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

	var err = http2.ConfigureServer(s.server, nil)
	if err != nil {
		return nil, err
	}

	return s, err
}

func (s *Server) Launch() {
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

func (s *Server) Stop() {
	var err error
	if err = s.server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
