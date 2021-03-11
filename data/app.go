package data

import "github.com/jackc/pgx/v4"

func NewApp() (*App, error) {
	var s App
	var err error

	// TODO: Data.App constructor

	return &s, err
}

type App struct {
	db *pgx.Conn
}
