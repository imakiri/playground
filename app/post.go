package app

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
)

type ConfigPost interface {
	Get4AppPost(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.AppPost, error)
}

type ServicePost struct {
	config       ConfigPost
	configCached *cfg.AppPost
}

// ServicePost.Create()
// ServicePost.Update()
// ServicePost.Delete()

func NewServicePost(c ConfigPost) (*ServicePost, error) {
	var s ServicePost
	var err error

	s.config = c
	s.configCached, err = s.config.Get4AppPost(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
