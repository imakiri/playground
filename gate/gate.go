package gate

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/transport"
)

type Service struct {
	log       core.LogService
	config    *transport.Gate
	configSys *transport.System
}

func NewService(c *transport.Config) (*Service, error) {
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
