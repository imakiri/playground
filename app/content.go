package app

import (
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/core"
	"github.com/jackc/pgx/v4"
)

type Service struct {
	db       *pgx.Conn
	log      core.LogService
	config   *cfg.Content
	configDB *cfg.Data
}

func NewContentService(c *cfg.System) (*Service, error) {
	var s Service
	var err error

	s.config = c.GetContent()
	s.configDB = c.GetData()

	s.db, err = core.Connect(c.GetData())
	if err != nil {
		return nil, err
	}

	return &s, err
}
