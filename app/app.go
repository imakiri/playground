package app

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/service"
)

type Data interface {
	SomeDataFunc()
}

type Service struct {
	service.Service
	config *cfg.App
	data   Data
}

func New(bs service.Service) (*Service, error) {
	var s Service
	var err error

	s.Service = bs
	s.config, err = s.Cfg().Get4App(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
