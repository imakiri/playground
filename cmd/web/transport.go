package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

func connect(o opts) (*grpc.ClientConn, error) {
	var ips, err = net.LookupIP(o.domain)
	if err != nil {
		return nil, err
	}

	var creds credentials.TransportCredentials
	creds, err = credentials.NewClientTLSFromFile(path_cert, o.domain)
	if err != nil {
		return nil, err
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(ips[0].String()+":"+o.port, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return conn, err
}
