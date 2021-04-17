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

const (
	MONGODB  = "mongodb"
	POSTGRES = "postgres"
)

/* Получается странная связанность, если попытаться обьединить 2 дб-методы в 1 сервис-конструктор.

   Необходимо будет принимать сразу 2 аргумента с разными дб, что нам тогда ставить? 1 nil в аргумент?
*/

func NewAvatarService(dbType string, db *mongo.Database) (Avatar, error) {
	switch dbType {
	case MONGODB:
		return mongodb.NewMongoRepos(db), nil
	case POSTGRES:
		return nil
	}
}
