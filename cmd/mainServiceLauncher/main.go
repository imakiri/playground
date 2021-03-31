package main

import (
	"context"
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/transport"
	"github.com/imakiri/gorum/types"
	"github.com/imakiri/gorum/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	bi_log "log"
	"net"
)

type opts struct {
	domain   string
	port     string
	certPath string
}

func NewLauncher(otps opts) (*Launcher, error) {
	var l Launcher
	var err error

	var ips []net.IP
	ips, err = net.LookupIP(otps.domain)
	if err != nil {
		return nil, err
	}

	var creds credentials.TransportCredentials
	creds, err = credentials.NewClientTLSFromFile(otps.certPath, otps.domain)
	if err != nil {
		return nil, err
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(ips[0].String()+":"+otps.port, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	l.config = transport.NewConfigClient(conn)
	return &l, err
}

type Launcher struct {
	config transport.ConfigClient
}

func (l *Launcher) Launch() error {
	var config, err = l.config.Get4Web(context.Background(), &types.ConfigRequest{})
	if err != nil {
		return erres.InternalServiceError.Extend()
	}

	var w *web.Service
	w, err = web.NewService(config)
	if err != nil {
		return erres.InternalServiceError.Extend()
	}

	return w.Launch()
}

const (
	Domain   = "imakiri-ips.ddns.net"
	Port     = "25565"
	CertPath = "secrets/grpc/cert.crt"
)

func main() {
	var o opts

	o.domain = Domain
	o.port = Port
	o.certPath = CertPath

	var l, err = NewLauncher(o)
	if err != nil {
		bi_log.Fatalln(err)
	} else {
		bi_log.Fatalln(l.Launch())
	}
}
