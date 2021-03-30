package cfg

import (
	"context"
	"fmt"
	"github.com/imakiri/gorum/transport"
	"github.com/imakiri/gorum/types"
	"github.com/spf13/viper"
)

func New() (*Service, error) {
	var s Service
	var err error

	s.config, err = readConfig()
	if err != nil {
		return nil, err
	}

	return &s, err
}

type Service struct {
	transport.UnimplementedConfigServer
	config *types.Config
}

func (e Service) Get4Api(_ context.Context, _ *types.ConfigRequest) (*types.ConfigApi, error) {
	fmt.Println("config sent")
	fmt.Println(e.config.Api.String())
	return e.config.Api, nil
}

func (e Service) Get4App(_ context.Context, _ *types.ConfigRequest) (*types.ConfigApp, error) {
	fmt.Println("config sent")
	fmt.Println(e.config.App.String())
	return e.config.App, nil
}

func (e Service) Get4Auth(_ context.Context, _ *types.ConfigRequest) (*types.ConfigAuth, error) {
	fmt.Println("config sent")
	fmt.Println(e.config.Auth.String())
	return e.config.Auth, nil
}

func (e Service) Get4Data(_ context.Context, _ *types.ConfigRequest) (*types.ConfigData, error) {
	fmt.Println("config sent")
	fmt.Println(e.config.Data.String())
	return e.config.Data, nil
}

func (e Service) Get4DataPostgres(_ context.Context, _ *types.ConfigRequest) (*types.ConfigDataPostgres, error) {
	fmt.Println("config sent")
	fmt.Println(e.config.Data.Postgres.String())
	return e.config.Data.Postgres, nil
}

func (e Service) Get4DataMongo(_ context.Context, _ *types.ConfigRequest) (*types.ConfigDataMongo, error) {
	fmt.Println("config sent")
	fmt.Println(e.config.Data.Mongo.String())
	return e.config.Data.Mongo, nil
}

func (e Service) Get4Gate(_ context.Context, _ *types.ConfigRequest) (*types.ConfigGate, error) {
	fmt.Println("config sent")
	fmt.Println(e.config.Gate.String())
	return e.config.Gate, nil
}

func (e Service) Get4Web(_ context.Context, _ *types.ConfigRequest) (*types.ConfigWeb, error) {
	fmt.Println("config sent")
	fmt.Println(e.config.Web.String())
	return e.config.Web, nil
}

func readConfig() (*types.Config, error) {
	var conf = new(types.Config)
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

	return conf, err
}
