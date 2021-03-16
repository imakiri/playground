package data

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/service"
)

type Auth struct {
	service.Service
	config *cfg.DataAuth
}

func NewAuth(bs service.Service) (*Auth, error) {
	var s Auth
	var err error

	s.Service = bs
	s.config, err = s.Cfg().Get4DataAuth(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	// TODO: Auth constructor

	return &s, err
}
