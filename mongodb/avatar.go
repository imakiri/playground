package mongodb

import (
	"context"

	"github.com/imakiri/gorum/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Avatar struct {
	avatar []byte
}

func AvatarGet128(db *mongo.Collection, userUUID types.ModelUserUUID, container *types.ModelUserAvatar128) (*Avatar, error) {
	var data Avatar
	objId, err := primitive.ObjectIDFromHex(string(userUUID))
	if err != nil {
		return nil, err
	}
	err = db.FindOne(context.Background(), bson.M{"id": objId, "avatar128": container}).Decode(&data)
	return &data, err
}

func AvatarGet256(db *mongo.Collection, userUUID types.ModelUserUUID, container *types.ModelUserAvatar256) (*Avatar, error) {
	var data Avatar
	objId, err := primitive.ObjectIDFromHex(string(userUUID))
	if err != nil {
		return nil, err
	}
	err = db.FindOne(context.Background(), bson.M{"id": objId, "avatar256": container}).Decode(&data)
	return &data, err
}

func AvatarGet512(db *mongo.Collection, userUUID types.ModelUserUUID, container *types.ModelUserAvatar512) (*Avatar, error) {
	var data Avatar
	objId, err := primitive.ObjectIDFromHex(string(userUUID))
	if err != nil {
		return nil, err
	}
	err = db.FindOne(context.Background(), bson.M{"id": objId, "avatar512": container}).Decode(&data)
	return &data, err
}

func AvatarSet(db *mongo.Collection, update bool, userUUID types.ModelUserUUID, avatar types.ViewUserAvatar) error {
	objId, err := primitive.ObjectIDFromHex(string(userUUID))
	if err != nil {
		return err
	}

	if update {
		res := db.FindOneAndUpdate(context.Background(), bson.M{"id": objId}, bson.M{
			"$set": bson.M{
				"avatar128": avatar.Avatar128,
				"avatar256": avatar.Avatar256,
				"avatar512": avatar.Avatar512,
			},
		})

		return res.Err()
	}

	_, err = db.InsertOne(context.Background(), bson.M{
		"id":        objId,
		"avatar128": avatar.Avatar128,
		"avatar256": avatar.Avatar256,
		"avatar512": avatar.Avatar512,
	})

	return err
}
