package main

import (
	"github.com/imakiri/gorum/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	bi_log "log"
	"net"
)

const path_cert = "secrets/grpc/cert.crt"

type opts struct {
	domain string
	port   string
}

func connect(otps opts) (*grpc.ClientConn, error) {
	var ips, err = net.LookupIP(otps.domain)
	if err != nil {
		return nil, err
	}

	var creds credentials.TransportCredentials
	creds, err = credentials.NewClientTLSFromFile(path_cert, otps.domain)
	if err != nil {
		return nil, err
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(ips[0].String()+":"+otps.port, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return conn, err
}

func NewLauncher(otps opts) (*Launcher, error) {
	var l Launcher
	var err error

	l.web, err = web.NewService()
	if err != nil {
		return nil, err
	}

	return &l, err
}

type Launcher struct {
	web *web.Service
}

func (l *Launcher) Launch() error {
	return l.web.Launch()
}

const (
	debug_domain = "imakiri-ips.ddns.net"
	default_port = "25565"
)

func main() {
	var o opts

	o.domain = debug_domain
	o.port = default_port

	var l, err = NewLauncher(o)
	if err != nil {
		bi_log.Fatalln(err)
	} else {
		bi_log.Fatalln(l.Launch())
	}
}
