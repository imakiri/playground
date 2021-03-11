package cfg

import (
	"context"
	"fmt"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/transport"
	"github.com/spf13/viper"
)

type Config struct {
	Api  *core.CfgApi
	App  *core.CfgApp
	Auth *core.CfgAuth
	Data *core.CfgData
	Gate *core.CfgGate
	Web  *core.CfgWeb
}

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
	transport.UnimplementedCfgServer
	config *Config
}

func (e Service) Api(_ context.Context, _ *core.Request) (*core.CfgApi, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Api.String())
	return e.config.Api, nil
}

func (e Service) App(_ context.Context, _ *core.Request) (*core.CfgApp, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.App.String())
	return e.config.App, nil
}

func (e Service) Auth(_ context.Context, _ *core.Request) (*core.CfgAuth, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Auth.String())
	return e.config.Auth, nil
}

func (e Service) Data(_ context.Context, _ *core.Request) (*core.CfgData, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Data.String())
	return e.config.Data, nil
}

func (e Service) Gate(_ context.Context, _ *core.Request) (*core.CfgGate, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.Gate.String())
	return e.config.Gate, nil
}

func (e Service) Web(_ context.Context, _ *core.Request) (*core.CfgWeb, error) {
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
