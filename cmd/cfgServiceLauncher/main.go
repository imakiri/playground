package main

import (
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type opts struct {
	port     string
	certPath string
	keyPath  string
}

func NewLauncher(o opts) (*Launcher, error) {
	var l Launcher
	var err error

	l.lis, err = net.Listen("tcp", ":"+o.port)
	if err != nil {
		return nil, err
	}

	var creds credentials.TransportCredentials
	creds, err = credentials.NewServerTLSFromFile(o.certPath, o.keyPath)
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

const (
	Port     = "25565"
	CertPath = "cfg/grpc/cert.crt"
	KeyPath  = "cfg/grpc/key.pem"
)

func main() {
	var o opts

	//o.port = string(*flag.Int("port", Port, "port of cfg server"))
	o.port = Port
	o.certPath = CertPath
	o.keyPath = KeyPath

	var l, err = NewLauncher(o)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Fatalln(l.Launch())
	}

}
