package app

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
)

type ConfigUser interface {
	Get4AppUser(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.AppUser, error)
}

type ServiceUser struct {
	config       ConfigUser
	configCached *cfg.AppUser
}

// ServiceUser.GetProfile()
// ServiceUser.UpdateProfile()

func NewServiceUser(c ConfigUser) (*ServiceUser, error) {
	var s ServiceUser
	var err error

	s.config = c
	s.configCached, err = s.config.Get4AppUser(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
