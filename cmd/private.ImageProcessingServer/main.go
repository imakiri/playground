package main

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/ips"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	gs := grpc.NewServer()

	ipsInstance := ips.NewFaceDetector()
	core.RegisterFaceDetecterServer(gs, ipsInstance)
	reflection.Register(gs)

	nl, err := net.Listen("tcp", ":25565")
	if err != nil {
		panic("Unable to bind to the port 25565")
	}

	_ = gs.Serve(nl)
}
