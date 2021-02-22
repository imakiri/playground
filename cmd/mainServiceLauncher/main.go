package main

import (
	"context"
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/ei/web"
	"github.com/imakiri/playground/gate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func readConfig() (*cfg.Config, error) {
	var conf cfg.Config
	var err error

	var ips []net.IP
	ips, err = net.LookupIP("imakiri-ips.ddns.net")
	if err != nil {
		log.Fatal(err)
	}

	var client cfg.ConfigClient
	client, err = launchConfigClient(ips[0].String()+":25565", "cfg/grpc/cert.crt")
	if err != nil {
		log.Fatal(err)
	}

	var data *cfg.Data
	data, err = client.RequestConfig(context.Background(), &cfg.Request{})
	if err != nil {
		log.Fatal(err)
	}

	conf.System.DB.DSN = data.DSN
	return &conf, err
}

func launchConfigClient(addr, certFile string) (cfg.ConfigClient, error) {
	var client cfg.ConfigClient
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

	client = cfg.NewConfigClient(conn)
	return client, err
}

func launch(c *cfg.Config) error {
	var err error
	var gs gate.GeneralService
	var ws *web.Service

	gs, err = gate.NewService(c)
	if err != nil {
		return err
	}

	ws, err = web.NewService(c.EI, gs)
	if err != nil {
		return err
	}

	rsc := make(chan error)
	sc := make(chan error)

	go func(rsc chan error) {
		rsc <- ws.RedirServer.ListenAndServe()
	}(rsc)

	go func(sc chan error) {
		sc <- ws.Server.ListenAndServeTLS("cfg/http/cert.pem", "cfg/http/privkey.pem")
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
