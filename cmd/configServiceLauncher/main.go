package main

import (
	"context"
	"fmt"
	"github.com/imakiri/playground/cfg"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type ConfigService struct {
	cfg.UnimplementedConfigServer
	config *cfg.Config
}

func (e ConfigService) RequestConfig(_ context.Context, _ *cfg.Request) (*cfg.Data, error) {
	var re cfg.Data
	var err error

	var ips []net.IP
	ips, err = net.LookupIP("imakiri-ips.ddns.net")
	if err != nil {
		return nil, err
	}

	re.DSN = "postgres://service:a5l6d99@" + ips[0].String() + ":5432/data?sslmode=disable"

	fmt.Println("Config sent")
	return &re, err
}

func readConfig() (*cfg.Config, error) {
	var conf cfg.Config
	var err error
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./cfg/")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, err
}

func launchService(c *cfg.Config, addr, certFile, keyFile string) error {
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

	var service = ConfigService{config: c}
	cfg.RegisterConfigServer(server, service)

	return server.Serve(lis)
}

func main() {
	var err error
	var conf *cfg.Config

	conf, err = readConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = launchService(conf, ":25565", "cfg/grpc/cert.crt", "cfg/grpc/key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
