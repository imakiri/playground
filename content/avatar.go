package content

import (
	"github.com/imakiri/gorum/postgres"
	"github.com/imakiri/gorum/types"
)

type avatarPostgres struct {
	connectionPostgres
}

func (s avatarPostgres) Get128(userUUID types.ModelUserUUID, container *types.ModelUserAvatar128) error {
	return postgres.AvatarGet128(s.db, userUUID, container)
}

func (s avatarPostgres) Get256(userUUID types.ModelUserUUID, container *types.ModelUserAvatar256) error {
	return postgres.AvatarGet256(s.db, userUUID, container)
}

func (s avatarPostgres) Get512(userUUID types.ModelUserUUID, container *types.ModelUserAvatar512) error {
	return postgres.AvatarGet512(s.db, userUUID, container)
}

func (s avatarPostgres) Set(update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error {
	return postgres.AvatarSet(s.db, update, userUUID, avatar)
}

type avatarMongo struct {
	connectionMongo
}

func (s avatarMongo) Get128(userUUID types.ModelUserUUID, container *types.ModelUserAvatar128) error {
	panic("")
}

func (s avatarMongo) Get256(userUUID types.ModelUserUUID, container *types.ModelUserAvatar256) error {
	panic("")
}

func (s avatarMongo) Get512(userUUID types.ModelUserUUID, container *types.ModelUserAvatar512) error {
	panic("")
}

func (s avatarMongo) Set(update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error {
	panic("")
}
