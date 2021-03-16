package main

import (
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type opts struct {
	port      string
	cert_path string
	key_path  string
}

func NewLauncher(o opts) (*Launcher, error) {
	var l Launcher
	var err error

	l.lis, err = net.Listen("tcp", o.port)
	if err != nil {
		return nil, err
	}

	var creds credentials.TransportCredentials
	creds, err = credentials.NewServerTLSFromFile(o.cert_path, o.key_path)
	if err != nil {
		return nil, err
	}

	l.server = grpc.NewServer(grpc.Creds(creds))

	var service *cfg.Service
	service, err = cfg.New()
	if err != nil {
		return nil, err
	}

	cfg.RegisterServiceServer(l.server, service)

	return &l, err
}

type Launcher struct {
	lis    net.Listener
	server *grpc.Server
}

func (l *Launcher) Launch() error {
	return l.server.Serve(l.lis)
}

func main() {

	// TODO: Grab func args from command line args

	var o opts
	o.port = ":25565"
	o.cert_path = "cfg/grpc/cert.crt"
	o.key_path = "cfg/grpc/key.pem"

	var l, err = NewLauncher(o)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Fatalln(l.Launch())
	}

}
