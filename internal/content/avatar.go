package content

import (
	"github.com/imakiri/erres"
	"github.com/imakiri/gorum/internal/auth"
	"github.com/imakiri/gorum/internal/cfg"
	"github.com/imakiri/gorum/internal/postgres"
)

type serviceAvatarPostgres struct {
	postgres.Connection
}

func (s serviceAvatarPostgres) Get128(userID auth.UserID, container []byte) error {
	return postgres.AvatarGet128(s.Connection, string(userID), container)
}

func (s serviceAvatarPostgres) Get256(userID auth.UserID, container []byte) error {
	return postgres.AvatarGet256(s.Connection, string(userID), container)
}

func (s serviceAvatarPostgres) Get512(userID auth.UserID, container []byte) error {
	return postgres.AvatarGet512(s.Connection, string(userID), container)
}

func (s serviceAvatarPostgres) Set(update bool, userID auth.UserID, avatar postgres.ViewUserAvatar) error {
	return postgres.AvatarSet(s.Connection, update, string(userID), avatar)
}

type serviceAvatarMongo struct {
	//mongodb.Connection
}

func (s serviceAvatarMongo) Get128(userID auth.UserID, container []byte) error {
	panic("implement me")
}

func (s serviceAvatarMongo) Get256(userID auth.UserID, container []byte) error {
	panic("implement me")
}

func (s serviceAvatarMongo) Get512(userID auth.UserID, container []byte) error {
	panic("implement me")
}

func (s serviceAvatarMongo) Set(update bool, userID auth.UserID, avatar postgres.ViewUserAvatar) error {
	panic("implement me")
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

		//var connection *mongodb.Connection
		//if connection, err = mongodb.NewConnection(nil, *config, *secret); err != nil {
		//	return nil, err
		//}

		var service = new(serviceAvatarMongo)
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

		var service = new(serviceAvatarPostgres)
		service.Connection = *connection

		return service, nil
	default:
		panic(erres.UnacceptableStateOfExecution)
	}
}
