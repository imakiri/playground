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

type connectionMongo struct {
	db *mongo.Database
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
	return s, nil
}

func newConnectionMongo() (*connectionMongo, error) {
	// cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password   <----- something like this we need in config
	mongoClient, err := mongodb.NewClient("", "", "")
	if err != nil {
		// fmt.Fatal(err)  <--- we need logger?
		return nil, erres.ConnectionError.Extend(0)
	}
	// something cg.Mongo.Name (name of db) in parameters we need
	db := mongoClient.Database("")

	return &connectionMongo{db: db}, nil
}

const (
	MONGODB  = "mongodb"
	POSTGRES = "postgres"
)

func AuthService(dbType string) error {
	switch dbType {
	case MONGODB:
		// cfg.mongo parameters we need soon for connection func
		mongo, err := newConnectionMongo()
		// here we need to provide db argument for our dirty postgres/mongo methods :D
		return err
	case POSTGRES:
		postgres, err := newConnectionPostgres()
		// here we need to provide db argument for our dirty postgres/mongo methods :D
		return err
	default:
		return nil // Where is default err of AuthSerivce? How to generate err with your library errers, ikamiri?
	}
}
