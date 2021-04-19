package api

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
)

type Config interface {
	Get4Api(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.Api, error)
}

type Service struct {
	config       Config
	configCached *cfg.Api
}

func New(c Config) (*Service, error) {
	var s Service
	var err error

	s.config = c
	s.configCached, err = s.config.Get4Api(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
