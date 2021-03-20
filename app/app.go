package app

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
)

type DataAuth interface{}
type DataApp interface{}

type Config interface {
	Get4App(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.App, error)
}

type Service struct {
	config       Config
	configCached *cfg.App
	dataAuth     DataAuth
	dataApp      DataApp
}

func New(c Config) (*Service, error) {
	var s Service
	var err error

	s.config = c
	s.configCached, err = s.config.Get4App(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
