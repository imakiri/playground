package app

import (
	"github.com/imakiri/gorum/core"
	"github.com/imakiri/gorum/data"
	"github.com/imakiri/gorum/transport"
	"github.com/jackc/pgx/v4"
)

type Service struct {
	db       *pgx.Conn
	log      core.LogService
	config   *transport.Content
	configDB *transport.Data
}

func NewContentService(c *transport.System) (*Service, error) {
	var s Service
	var err error

	s.config = c.GetContent()
	s.configDB = c.GetData()

	s.db, err = data.Connect(c.GetData())
	if err != nil {
		return nil, err
	}

	return &s, err
}
