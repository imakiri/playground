package content

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/types"
	"github.com/imakiri/gorum/utils"
	"github.com/jmoiron/sqlx"
)

const path = "secrets/content.yaml"

type service struct {
	secret types.SecretData
	db     *sqlx.DB
}

func newService() (*service, error) {
	var secret types.SecretData
	var err = utils.ReadYAML(path, &secret)
	if err != nil {
		return nil, erres.InternalServiceError.Extend(0)
	}

	var s service
	s.secret = secret
	if s.db, err = sqlx.Connect("pgx", secret.Postgres.DSN); err != nil {
		return nil, erres.ConnectionError.Extend(0)
	}
	return &s, err
}

func NewAvatarService() {

}
