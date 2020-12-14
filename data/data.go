package data

import (
	"github.com/imakiri/playground/core"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	core.Settings
	db *sqlx.DB
}

func NewDB(s core.Settings) (*DB, error) {
	var err error
	var data DB

	data.Settings = s
	data.db, err = sqlx.Connect("pgx", data.Settings.Config.DSN)
	if err != nil {
		return nil, err
	}

	err = data.db.Ping()
	if err != nil {
		return nil, err
	}

	return &data, nil
}
