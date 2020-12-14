package data

import (
	"github.com/imakiri/playground/core"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"
)

type Error string

func (e Error) Error() string {
	return string(e)
}

type DB struct {
	db *sqlx.DB
}

func NewDB(s core.Settings) (*DB, error) {
	var err error
	var data DB

	data.db, err = sqlx.Connect("pgx", s.Config.Data.DSN)
	if err != nil {
		return nil, err
	}

	err = data.db.Ping()
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type WebClient struct {
	apiKey string
}

func NewWebClient(s core.Settings) (*WebClient, error) {
	var wc WebClient
	wc.apiKey = s.Config.Data.ApiKey
	return &wc, nil
}
