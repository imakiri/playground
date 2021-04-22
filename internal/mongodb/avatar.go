package mongodb

import (
	"github.com/imakiri/gorum/internal/data"
	"go.mongodb.org/mongo-driver/mongo"
)

func AvatarGet128(db *mongo.Database, userUUID data.ModelUserUUID, container *data.ModelUserAvatar128) error {
	return nil
}
func AvatarGet256(db *mongo.Database, userUUID data.ModelUserUUID, container *data.ModelUserAvatar256) error {
	return nil
}
func AvatarGet512(db *mongo.Database, userUUID data.ModelUserUUID, container *data.ModelUserAvatar512) error {
	return nil
}
func AvatarSet(db *mongo.Database, update bool, userUUID data.ModelUserUUID, avatar data.ViewUserAvatar) error {
	return nil
}
