package gate

import (
	"context"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/log"
)

func NewService(cfg_sc cfg.ServiceClient) (*Service, error) {
	var s Service
	var err error

	s.cfg_sc = cfg_sc
	s.config, err = s.cfg_sc.Get4Gate(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	return &s, err
}

type Service struct {
	log    log.Service
	cfg_sc cfg.ServiceClient
	config *cfg.Gate
}
