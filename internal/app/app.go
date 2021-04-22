package app

import (
	"context"
	"github.com/imakiri/gorum/internal/cfg"
)

type Service struct {
	config       Config
	configCached *cfg.App
	dataAuth     DataAuth
	dataApp      DataApp
}

func NewService(c Config) (*Service, error) {
	var s Service
	var err error

	s.config = c
	s.configCached, err = s.config.Get4App(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
