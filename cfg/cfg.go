package cfg

import (
	"context"
	"fmt"
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
	UnimplementedServiceServer
	config *Config
}

func (e Service) Get4Api(_ context.Context, _ *Request) (*Api, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Api.String())
	return e.config.Api, nil
}

func (e Service) Get4App(_ context.Context, _ *Request) (*App, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.App.String())
	return e.config.App, nil
}

func (e Service) Get4Auth(_ context.Context, _ *Request) (*Auth, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Auth.String())
	return e.config.Auth, nil
}

func (e Service) Get4Data(_ context.Context, _ *Request) (*Data, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Data.String())
	return e.config.Data, nil
}

func (e Service) Get4DataApp(context.Context, *Request) (*DataApp, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Data.App.String())
	return e.config.Data.App, nil
}

func (e Service) Get4DataAuth(context.Context, *Request) (*DataAuth, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Data.App.String())
	return e.config.Data.Auth, nil
}

func (e Service) Get4DataGate(context.Context, *Request) (*DataGate, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Data.App.String())
	return e.config.Data.Gate, nil
}

func (e Service) Get4Gate(_ context.Context, _ *Request) (*Gate, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Gate.String())
	return e.config.Gate, nil
}

func (e Service) Get4Web(_ context.Context, _ *Request) (*Web, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Web.String())
	return e.config.Web, nil
}

func readConfig() (*Config, error) {
	var conf = new(Config)
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
