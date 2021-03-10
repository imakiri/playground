package app

import (
	"github.com/imakiri/playground/cfg"
	"github.com/imakiri/playground/core"
	"github.com/jackc/pgx/v4"
)

type User struct {
	db       *pgx.Conn
	log      core.LogService
	config   *cfg.App
	configDB *cfg.Data
}

func NewService(c *cfg.System) (*User, error) {
	var s User
	var err error

	s.config = c.GetApp()
	s.configDB = c.GetData()

	s.db, err = core.Connect(c.GetData())
	if err != nil {
		return nil, err
	}

	return &s, err
}
