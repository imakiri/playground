package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Avatar *AvatarMongo
}

func NewMongoRepos(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Avatar: NewAvatarMongo(db),
	}
}
