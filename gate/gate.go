package gate

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
)

type Config interface {
	Get4Gate(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.Gate, error)
}

type Service struct {
	config       Config
	configCached *cfg.Gate
}

func NewService(c Config) (*Service, error) {
	var s Service
	var err error

	s.config = c
	s.configCached, err = s.config.Get4Gate(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}

func (s Service) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	var f grpc.UnaryServerInterceptor

	// TODO: service.UnaryServerInterceptor

	return f
}
