package main

import (
	"github.com/imakiri/playground/transport"
	"github.com/imakiri/playground/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func newConfigClient(from string, port string, certFile string) (transport.CfgClient, error) {
	var client transport.CfgClient
	var err error

	var ips []net.IP
	ips, err = net.LookupIP(from)
	if err != nil {
		return nil, err
	}

	var creds credentials.TransportCredentials
	creds, err = credentials.NewClientTLSFromFile(certFile, "imakiri-ips.ddns.net")
	if err != nil {
		return nil, err
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(ips[0].String()+port, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	client = transport.NewCfgClient(conn)
	return client, err
}

func main() {
	var cc transport.CfgClient
	var err error

	// TODO: Grab func args from command line args

	cc, err = newConfigClient("imakiri-ips.ddns.net", ":25565", "cfg/grpc/cert.crt")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(web.NewService(cc))
}
