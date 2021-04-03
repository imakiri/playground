package web

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
)

const (
	path_key  = "secrets/web/key.txt"
	path_cert = "secrets/web/certificate.crt"
)

type registrar func(*http.Server) error

type Service struct {
	https  bool
	server *http.Server
	status chan error
}

func NewService(status chan error, https bool, reg registrar) (*Service, error) {
	var s Service
	s.https = https
	s.server = &http.Server{}
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

	var err = reg(s.server)
	if err != nil {
		return nil, err
	}

	return &s, err
}

func (s *Service) Launch() {
	var l net.Listener
	var err error

	if s.https {
		l, err = tls.Listen("tcp", ":443", s.server.TLSConfig)
		if err != nil {
			s.status <- err
		}
	} else {
		l, err = net.Listen("tcp", ":80")
		if err != nil {
			s.status <- err
		}
	}

	go func() {
		s.status <- s.server.Serve(l)
	}()
}

func (s *Service) Stop() {
	var err = s.server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
