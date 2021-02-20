package app

import (
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"github.com/jackc/pgx/v4"
)

type Service struct {
	db       *pgx.Conn
	log      core.LogService
	config   cfg.App
	configDB cfg.DB
}

func NewService(c cfg.System) (*Service, error) {
	var s Service
	var err error

	s.config = c.App
	s.configDB = c.DB

	s.db, err = data.Connect(c.DB)
	if err != nil {
		return nil, err
	}

	return &s, err
}
