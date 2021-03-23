package app

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"google.golang.org/grpc"
)

type ConfigThread interface {
	Get4AppThread(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.AppThread, error)
}

type ServiceThread struct {
	config       ConfigThread
	configCached *cfg.AppThread
}

// ServiceThread.Create()
// ServiceThread.GetThreadContent()
// ServiceThread.GetThreadsList()
// ServiceThread.Update()
// ServiceThread.Delete()

func NewServiceThread(c ConfigThread) (*ServiceThread, error) {
	var s ServiceThread
	var err error

	s.config = c
	s.configCached, err = s.config.Get4AppThread(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}
