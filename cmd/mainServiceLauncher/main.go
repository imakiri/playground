package main

import (
	"context"
	"github.com/imakiri/playground/transport"
	"github.com/imakiri/playground/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func getConfig(from string, port string, certFile string) (*transport.Config, error) {
	var conf *transport.Config
	var err error

	var ips []net.IP
	ips, err = net.LookupIP(from)
	if err != nil {
		log.Fatal(err)
	}

	var client transport.AdminClient
	client, err = launchConfigClient(ips[0].String()+port, certFile)
	if err != nil {
		log.Fatal(err)
	}

	conf, err = client.GetConfig(context.Background(), &transport.Request{})
	if err != nil {
		log.Fatal(err)
	}

	return conf, err
}

func launchConfigClient(addr, certFile string) (transport.AdminClient, error) {
	var client transport.AdminClient
	var err error

	var creds credentials.TransportCredentials
	creds, err = credentials.NewClientTLSFromFile(certFile, "imakiri-ips.ddns.net")
	if err != nil {
		return nil, err
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	client = transport.NewAdminClient(conn)
	return client, err
}

func launch(c *transport.Config, certFile string, keyFile string) error {
	var err error
	var ws *web.Service

	ws, err = web.NewService(c.GetEI())
	if err != nil {
		return err
	}

	rsc := make(chan error)
	sc := make(chan error)

	go func(rsc chan error) {
		rsc <- ws.RedirServer.ListenAndServe()
	}(rsc)

	go func(sc chan error) {
		sc <- ws.Server.ListenAndServeTLS(certFile, keyFile)
	}(sc)

	select {
	case err := <-rsc:
		return err
	case err := <-sc:
		return err
	}
}

func main() {
	var conf *transport.Config
	var err error

	// TODO: Grab func args from command line args

	conf, err = getConfig("imakiri-ips.ddns.net", ":25565", "cfg/grpc/cert.crt")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(launch(conf, "cfg/web/certificate.crt", "cfg/web/key.txt"))
}
