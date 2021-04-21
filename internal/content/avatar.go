package content

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/cfg"
	"github.com/imakiri/gorum/internal/postgres"
	"github.com/imakiri/gorum/internal/types"
)

type avatarPostgres struct {
	postgres.Connection
}

func (s avatarPostgres) Get128(userUUID types.ModelUserUUID, container types.ModelUserAvatar128) error {
	return postgres.AvatarGet128(s.Connection, string(userUUID), container)
}

func (s avatarPostgres) Get256(userUUID types.ModelUserUUID, container types.ModelUserAvatar256) error {
	return postgres.AvatarGet256(s.Connection, string(userUUID), container)
}

func (s avatarPostgres) Get512(userUUID types.ModelUserUUID, container types.ModelUserAvatar512) error {
	return postgres.AvatarGet512(s.Connection, string(userUUID), container)
}

func (s avatarPostgres) Set(update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error {
	return postgres.AvatarSet(s.Connection, update, string(userUUID), avatar)
}

type avatarMongo struct {
	//connectionMongo
}

func (s avatarMongo) Get128(userUUID types.ModelUserUUID, container types.ModelUserAvatar128) error {
	panic("")
}

func (s avatarMongo) Get256(userUUID types.ModelUserUUID, container types.ModelUserAvatar256) error {
	panic("")
}

func (s avatarMongo) Get512(userUUID types.ModelUserUUID, container types.ModelUserAvatar512) error {
	panic("")
}

func (s avatarMongo) Set(update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error {
	panic("")
}

func NewServiceAvatar(cs *cfg.Service) (ServiceAvatar, error) {
	var err error
	var sconf *cfg.ConfigService
	if sconf, err = cs.GetAvatarConfig(); err != nil {
		return nil, err
	}

	switch sconf.DBType {
	case cfg.MONGO:
		//var config *cfg.ConfigDatabaseMongo
		//var secret *cfg.SecretDatabaseMongo
		//if config, err = cs.GetMongoConfig(cfgConnName); err != nil {
		//	return nil, err
		//}
		//if secret, err = cs.GetMongoSecret(cfgConnName); err != nil {
		//	return nil, err
		//}

		//var connection *postgres.Connection
		//if connection, err = postgres.NewConnection(nil, *config, *secret); err != nil {
		//	return nil, err
		//}

		var service = new(avatarMongo)
		//service.Connection = *connection

		return service, nil
	case cfg.POSTGRES:
		var config *cfg.ConfigDatabasePostgres
		var secret *cfg.SecretDatabasePostgres
		if config, err = cs.GetPostgresConfig(sconf.ConnName); err != nil {
			return nil, err
		}
		if secret, err = cs.GetPostgresSecret(sconf.ConnName); err != nil {
			return nil, err
		}

		var connection *postgres.Connection
		if connection, err = postgres.NewConnection(nil, *config, *secret); err != nil {
			return nil, err
		}

		var service = new(avatarPostgres)
		service.Connection = *connection

		return service, nil
	default:
		panic(erres.UnacceptableStateOfExecution)
	}
}
