package content

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/mongodb"
	"github.com/imakiri/gorum/postgres"
	"github.com/imakiri/gorum/types"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

const path = "secrets/content.yaml"

type typeDB string

const (
	MONGODB  typeDB = "mongodb"
	POSTGRES typeDB = "postgres"
)

type connectionPostgres struct {
	secret *cfg.SecretDatabasePostgres
	cfg    *cfg.ConfigDatabasePostgres
	db     *sqlx.DB
}

type connectionMongo struct {
	cfg    *cfg.ConfigDatabaseMongo
	secret *cfg.SecretDatabaseMongo
	db     *mongo.Database
}

//func newConnectionMongo() (*connectionMongo, error) {
//	// cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password   <----- something like this we need in config
//	mongoClient, err := mongodb.NewClient("", "", "")
//	if err != nil {
//		// fmt.Fatal(err)  <--- we need logger?
//		return nil, erres.ConnectionError.Extend(0)
//	}
//	// something cg.Mongo.Name (name of db) in parameters we need
//	db := mongoClient.Database("")
//
//	return &connectionMongo{db: db}, nil
//}

type ServiceAvatar interface {
	Get128(userUUID types.ModelUserUUID, container *types.ModelUserAvatar128) error
	Get256(userUUID types.ModelUserUUID, container *types.ModelUserAvatar256) error
	Get512(userUUID types.ModelUserUUID, container *types.ModelUserAvatar512) error
	Set(update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error
}

func NewServiceAvatar(cs *cfg.Service, dbType typeDB, cfgType cfg.TypeCfg) (ServiceAvatar, error) {
	var err error

	switch dbType {
	case MONGODB:
		var service = new(avatarMongo)
		if service.cfg, err = cs.GetMongoCfg(cfgType); err != nil {
			return nil, err
		}
		if service.secret, err = cs.GetMongoSecret(cfgType); err != nil {
			return nil, err
		}
		if service.db, err = mongodb.NewClient(service.secret.DSN); err != nil {
			return nil, err
		}
		return service, nil
	case POSTGRES:
		var service = new(avatarPostgres)
		if service.cfg, err = cs.GetPostgresCfg(cfgType); err != nil {
			return nil, err
		}
		if service.secret, err = cs.GetPostgresSecret(cfgType); err != nil {
			return nil, err
		}
		if service.db, err = postgres.NewDB(service.secret.DSN); err != nil {
			return nil, err
		}
		return service, nil
	default:
		panic(erres.UnacceptableStateOfExecution)
	}
}
