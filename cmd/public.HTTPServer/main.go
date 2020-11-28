package main

import (
	"github.com/imakiri/playground/protos"
	"github.com/imakiri/playground/web"
	"google.golang.org/grpc"
	"log"
)

var gc protos.FaceDetecterClient

func main() {
	gsConn, err := grpc.Dial("localhost:25565", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer gsConn.Close()

	gc = protos.NewFaceDetecterClient(gsConn)

	log.Fatal(web.NewWebServer(gc))
}
