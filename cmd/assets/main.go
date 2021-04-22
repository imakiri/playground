package main

import (
	"github.com/imakiri/gorum/internal/app"
	"log"
	"net"
)

const (
	path_cert = "secrets/grpc/cert.crt"
	path_key  = "secrets/grpc/key.pem"
)

type opts struct {
	port     string
	certPath string
	keyPath  string
}

func NewLauncher(o opts) (*Launcher, error) {
	var l Launcher
	var err error

	//l.lis, err = net.Listen("tcp", ":"+o.port)
	//if err != nil {
	//	return nil, err
	//}
	//
	//var creds credentials.TransportCredentials
	//creds, err = credentials.NewServerTLSFromFile(o.certPath, o.keyPath)
	//if err != nil {
	//	return nil, err
	//}

	return &l, err
}

type Launcher struct {
	lis net.Listener
	app app.Service
}

func (l *Launcher) Launch() error {
	return nil
}

const (
	default_port = "25565"
)

func main() {
	var o opts

	//o.port = string(*flag.Int("port", default_port, "port of cfg server"))
	o.port = default_port
	o.certPath = path_cert
	o.keyPath = path_key

	var l, err = NewLauncher(o)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Fatalln(l.Launch())
	}

}
