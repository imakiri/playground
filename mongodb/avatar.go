package mongodb

import (
	"github.com/imakiri/gorum/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func AvatarGet128(db *mongo.Database, userUUID types.ModelUserUUID, container *types.ModelUserAvatar128) error {
	return nil
}
func AvatarGet256(db *mongo.Database, userUUID types.ModelUserUUID, container *types.ModelUserAvatar256) error {
	return nil
}
func AvatarGet512(db *mongo.Database, userUUID types.ModelUserUUID, container *types.ModelUserAvatar512) error {
	return nil
}
func AvatarSet(db *mongo.Database, update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error {
	return nil
}
