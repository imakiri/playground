package main

import (
	"flag"
	"github.com/imakiri/gorum/internal/web"
	"github.com/imakiri/gorum/internal/web/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	path_cert = "secrets/content/server.crt"
	path_key  = "secrets/content/server.key"
)

func NewLauncher(certPath string, keyPath string) (*Launcher, error) {
	var l = new(Launcher)

	var creds, err = credentials.NewServerTLSFromFile(certPath, keyPath)
	if err != nil {
		return nil, err
	}
	l.server = grpc.NewServer(grpc.Creds(creds))

	var service transport.ContentServer
	if service, err = web.NewContentService(); err != nil {
		return nil, err
	}
	transport.RegisterContentServer(l.server, service)

	return l, err
}

type Launcher struct {
	server *grpc.Server
}

func (l *Launcher) Launch(port string) error {
	var lis, err = net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	return l.server.Serve(lis)
}

const (
	port = "25565"
)

func main() {
	var port = flag.String("port", port, "port of content server")
	flag.Parse()

	var l, err = NewLauncher(path_cert, path_key)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(l.Launch(*port))
}
