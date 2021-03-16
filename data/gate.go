package data

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/service"
)

type Gate struct {
	service.Service
	config *cfg.DataGate
}

func New(bs service.Service) (*Gate, error) {
	var s Gate
	var err error

	s.Service = bs
	s.config, err = s.Cfg().Get4DataGate(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
