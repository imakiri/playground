package gate

import (
	"errors"
	"github.com/imakiri/playground/admin/cfg"
	"github.com/imakiri/playground/backend/app"
	"github.com/imakiri/playground/backend/content"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/gate/auth"
)

type Service struct {
	app       core.AppService
	auth      core.AuthService
	content   core.ContentService
	log       core.LogService
	config    *cfg.Gate
	configSys *cfg.System
}

func NewService(c *cfg.Config) (*Service, error) {
	var s Service
	var err error

	s.config = c.GetGate()
	s.configSys = c.GetSystem()

	s.app, err = app.NewService(c.GetSystem())
	if err != nil {
		return nil, err
	}

	s.auth, err = auth.NewService(c.GetSystem())
	if err != nil {
		return nil, err
	}

	s.content, err = content.NewContentService(c.GetSystem())
	if err != nil {
		return nil, err
	}

	return &s, err
}

func (e *Service) checkWorker() error {
	if e.app == nil {
		return errors.New("app error")
	}

	return nil
}
