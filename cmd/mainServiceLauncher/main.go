package main

import (
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/service"
	"github.com/imakiri/gorum/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	bi_log "log"
	"net"
)

type opts struct {
	domain    string
	port      string
	cert_path string
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
	creds, err = credentials.NewClientTLSFromFile(otps.cert_path, "imakiri-ips.ddns.net")
	if err != nil {
		return nil, err
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(ips[0].String()+otps.port, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	var config = cfg.NewServiceClient(conn)
	l.web, err = web.New(config)
	if err != nil {
		return nil, err
	}

	return &l, err
}

type Launcher struct {
	bs  service.Service
	web *web.Service
}

func (l *Launcher) Launch() error {
	return l.web.Launch()
}

func main() {

	// TODO: Grab func args from command line args

	var o opts
	o.domain = "imakiri-ips.ddns.net"
	o.port = ":25565"
	o.cert_path = "cfg/grpc/cert_path.crt"

	var l, err = NewLauncher(o)
	if err != nil {
		bi_log.Fatalln(err)
	} else {
		bi_log.Fatalln(l.Launch())
	}
}
