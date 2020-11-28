package main

import (
	"github.com/imakiri/playground/protos"
	"github.com/imakiri/playground/web"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

var conf Config
var gc protos.FaceDetecterClient

type Config struct {
	Database struct {
		User     string
		Password string
		Address  string
		Port     string
	}
	DSN       string
	ApiKey    string
	Salt      string
	IPSDomain string
}

func main() {
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

	gc = protos.NewFaceDetecterClient(gsConn)

	log.Fatal(web.NewWebServer(gc))
}
