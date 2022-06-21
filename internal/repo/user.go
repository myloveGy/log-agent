package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log-agent/internal/model"
	"log-agent/internal/util"
)

type UserRepo struct {
	Database *mongo.Database
}

func (u *UserRepo) Count(ctx context.Context, user *model.User) (int64, error) {
	return u.Database.Collection(user.Collection()).CountDocuments(ctx, util.StructNotZero2BsonM(user))
}

func (u *UserRepo) GetUser(ctx context.Context, username string) (*model.User, error) {
	item := &model.User{}
	if err := u.Database.Collection(item.Collection()).FindOne(ctx, bson.M{
		"username": username,
	}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (u *UserRepo) Create(ctx context.Context, user *model.User) error {
	if _, err := u.Database.Collection(user.Collection()).InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) Update(ctx context.Context, user *model.User) error {
	if _, err := u.Database.Collection(user.Collection()).UpdateOne(ctx, bson.M{
		"_id": user.Id,
	}, bson.D{
		{"$set", util.StructNotZero2BsonM(user)},
	}); err != nil {
		return err
	}

	return nil
}
