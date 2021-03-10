package main

import (
	"context"
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func readConfig() (*cfg.Config, error) {
	var conf *cfg.Config
	var err error

	var ips []net.IP
	ips, err = net.LookupIP("imakiri-ips.ddns.net")
	if err != nil {
		log.Fatal(err)
	}

	var client cfg.AdminClient
	client, err = launchConfigClient(ips[0].String()+":25565", "cfg/grpc/cert.crt")
	if err != nil {
		log.Fatal(err)
	}

	conf, err = client.GetConfig(context.Background(), &cfg.Request{})
	if err != nil {
		log.Fatal(err)
	}

	return conf, err
}

func launchConfigClient(addr, certFile string) (cfg.AdminClient, error) {
	var client cfg.AdminClient
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

	client = cfg.NewAdminClient(conn)
	return client, err
}

func launch(c *cfg.Config) error {
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
		sc <- ws.Server.ListenAndServeTLS("cfg/certificate.crt", "cfg/key.txt")
	}(sc)

	select {
	case err := <-rsc:
		return err
	case err := <-sc:
		return err
	}
}

func main() {
	var conf *cfg.Config
	var err error

	conf, err = readConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(launch(conf))
}
