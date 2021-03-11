package data

import (
	"github.com/jackc/pgx/v4"
)

func NewAuth() (*Auth, error) {
	var s Auth
	var err error

	// TODO: Data.Auth constructor

	return &s, err
}

type Auth struct {
	db *pgx.Conn
}
