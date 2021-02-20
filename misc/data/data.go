package data

import (
	"github.com/imakiri/playground/core"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type Data struct {
	db     *sqlx.DB
	log    core.LogService
	config core.ConfigDB
}

func NewDataService(s core.Settings) (*Data, error) {
	var data Data
	var err error

	data.config = s.Config.DB

	data.db, err = sqlx.Connect("pgx", data.config.DSN)
	if err != nil {
		return nil, err
	}

	err = data.db.Ping()
	if err != nil {
		return nil, err
	}

	return &data, err
}
