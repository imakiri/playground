package content

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/mongodb"
	"github.com/imakiri/gorum/types"
	"github.com/imakiri/gorum/utils"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

const path = "secrets/content.yaml"

type connectionPostgres struct {
	secret types.SecretData
	db     *sqlx.DB
}

func newConnectionPostgres() (*connectionPostgres, error) {
	var secret types.SecretData
	var err = utils.ReadYAML(path, &secret)
	if err != nil {
		return nil, erres.InternalServiceError.Extend(0)
	}

	var s = new(connectionPostgres)
	s.secret = secret
	if s.db, err = sqlx.Connect("pgx", secret.Postgres.DSN); err != nil {
		return nil, erres.ConnectionError.Extend(0)
	}