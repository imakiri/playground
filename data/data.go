package data

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/types"
	"github.com/imakiri/gorum/utils"
	"github.com/jmoiron/sqlx"
)

const path = "secrets/postgres.yaml"

type Service struct {
	secretPostgres *types.SecretPostgres
	db             *sqlx.DB
}

func NewService() (*Service, error) {
	var err error

	var config types.SecretPostgres
	err = utils.ReadYAML(path, &config)
	if err != nil {
		return nil, erres.InternalServiceError.Extend().AddDescription(err.Error())
	}

	var s Service
	s.db, err = sqlx.Connect("pgx", config.DSN)
	if err != nil {
		return nil, erres.ConnectionError.Extend().AddDescription(err.Error())
	}
	return &s, err
}
