package data

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/service"
)

type App struct {
	service.Service
	config *cfg.DataApp
}

func NewApp(bs service.Service) (*App, error) {
	var s App
	var err error

	s.Service = bs
	s.config, err = s.Cfg().Get4DataApp(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	// TODO: App constructor

	return &s, err
}
