package main

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/web"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func startWeb(s core.Settings) error {
	var server = &http.Server{}
	var redirServer = &http.Server{}

	router, redirRouter, _ := web.NewWebRouters(s)
	server.Handler = router
	redirServer.Handler = redirRouter

	rsc := make(chan error)
	sc := make(chan error)

	go func(rsc chan error) {
		rsc <- redirServer.ListenAndServe()
	}(rsc)

	go func(sc chan error) {
		sc <- server.ListenAndServeTLS("cert.pem", "privkey.pem")
	}(sc)

	select {
	case err := <-rsc:
		return err
	case err := <-sc:
		return err
	}
}

func main() {
	var conf core.Config
	var err error
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err)
	}

	ips_ip, err := net.LookupIP(conf.IPSDomain)
	if err != nil {
		log.Fatal(err)
	}

	gsConn, err := grpc.Dial(ips_ip[0].String()+":25565", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer gsConn.Close()

	gc := core.NewFaceDetecterClient(gsConn)
	s := core.Settings{
		Config:   conf,
		Services: core.Services{FaceDetection: gc},
	}

	log.Fatal(startWeb(s))
}
