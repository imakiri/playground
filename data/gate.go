package data

import "github.com/jackc/pgx/v4"

func NewGate() (*Gate, error) {
	var s Gate
	var err error

	// TODO: Data.Gate constructor

	return &s, err
}

type Gate struct {
	db *pgx.Conn
}
