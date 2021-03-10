package gate

import (
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/core"
)

type Service struct {
	log       core.LogService
	config    *cfg.Gate
	configSys *cfg.System
}

func NewService(c *cfg.Config) (*Service, error) {
	var s Service
	var err error

	s.config = c.GetGate()
	s.configSys = c.GetSystem()

	return &s, err
}

//func (e *Service) checkWorker() error {
//	if e.app == nil {
//		return errors.New("app error")
//	}
//
//	return nil
//}
