package cfg

import (
	"context"
	"fmt"
	"github.com/imakiri/playground/transport"
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
	transport.UnimplementedAdminServer
	config *transport.Config
}

func (e Service) GetConfig(_ context.Context, _ *transport.Request) (*transport.Config, error) {
	fmt.Println("Config sent")
	fmt.Println(e.config.String())
	return e.config, nil
}

func readConfig() (*transport.Config, error) {
	var conf transport.Config
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
