package main

import (
	"context"
	"fmt"
	"github.com/imakiri/playground/admin/cfg"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type AdminService struct {
	cfg.UnimplementedAdminServer
	config *cfg.Config
}

func (e AdminService) GetConfig(_ context.Context, _ *cfg.Request) (*cfg.Config, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.String())
	return e.config, nil
}

func readConfig() (*cfg.Config, error) {
	var conf cfg.Config
	var err error
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./admin/cfg/")

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

	var service = AdminService{config: c}
	cfg.RegisterAdminServer(server, service)

	return server.Serve(lis)
}

func main() {
	var err error
	var conf *cfg.Config

	conf, err = readConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = launchService(conf, ":25565", "admin/cfg/grpc/cert.crt", "admin/cfg/grpc/key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
