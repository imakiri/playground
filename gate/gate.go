package gate

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/misc"
)

type Service struct {
	app    core.ContentService
	auth   core.AuthService
	data   misc.DataService
	log    core.LogService
	config core.ConfigGate
}

func NewService(s core.Settings) (*Service, error) {
	var g Service
	var err error

	g.config = s.Config.Gate

	return &g, err
}
