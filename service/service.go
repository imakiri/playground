package service

import (
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/log"
)

func IsNilSafe(l ...interface{}) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			return false
		}
	}
	return true
}

func IsNilSafeEx(l ...interface{}) (b []bool) {
	for i := 0; i < len(l); i++ {
		if l[i] == nil {
			b = append(b, false)
		} else {
			b = append(b, true)
		}
	}
	return
}

type Service interface {
	Log() *log.Service
	Cfg() cfg.ServiceClient
}

func New(log *log.Service, cfgc cfg.ServiceClient) (*service, error) {
	var s service
	var err error

	s.log = log
	s.cfg = cfgc

	return &s, err
}

type service struct {
	log *log.Service
	cfg cfg.ServiceClient
}

func (s service) Log() *log.Service {
	return s.log
}

func (s service) Cfg() cfg.ServiceClient {
	return s.cfg
}
