package main

import (
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func launchService(addr, certFile, keyFile string) error {
	var err error

	var lis net.Listener
	lis, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var creds credentials.TransportCredentials
	creds, err = credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return err
	}

	var server *grpc.Server
	server = grpc.NewServer(grpc.Creds(creds))

	var service *cfg.Service
	service, err = cfg.New()
	if err != nil {
		return err
	}

	transport.RegisterAdminServer(server, service)
	return server.Serve(lis)
}

func main() {

	// TODO: Grab func args from command line args

	var err error
	err = launchService(":25565", "cfg/grpc/cert.crt", "cfg/grpc/key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
