package content

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"github.com/jackc/pgx/v4"
)

type Service struct {
	db       *pgx.Conn
	log      core.LogService
	config   core.ConfigContent
	configDB core.ConfigDB
}

func NewContentService(c core.Config) (*Service, error) {
	var s Service
	var err error

	s.config = c.Content
	s.configDB = c.DB

	s.db, err = data.Connect(c.DB)
	if err != nil {
		return nil, err
	}

	return &s, err
}
