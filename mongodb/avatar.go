package mongodb

import (
	"github.com/imakiri/gorum/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type AvatarMongo struct {
	db *mongo.Collection
}

func NewAvatarMongo(db *mongo.Database) *AvatarMongo {
	return &AvatarMongo{
		db: db.Collection("avatar-blah-blah"),
	}
}

func (av *AvatarMongo) Get128(userUUID types.ModelUserUUID, container *types.ModelUserAvatar128) error {
	return nil
}

func (av *AvatarMongo) Get256(userUUID types.ModelUserUUID, container *types.ModelUserAvatar256) error {
	return nil
}

func (av *AvatarMongo) Get512(userUUID types.ModelUserUUID, container *types.ModelUserAvatar512) error {
	return nil
}
