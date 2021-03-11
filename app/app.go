package app

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/transport"
	"github.com/jackc/pgx/v4"
)

type User struct {
	db       *pgx.Conn
	log      core.LogService
	config   *transport.App
	configDB *transport.Data
}

func NewService(c *transport.System) (*User, error) {
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
