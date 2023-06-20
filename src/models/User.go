package models

import (
	"context"
	"htmx-cares/src/core"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

func NewUserModel(email string, password string) (*UserModel) {
	return &UserModel{
		Email: email,
		Password: password,
	}
}

func (model *UserModel) Insert(collection *mongo.Collection) (*core.HttpErr) {
	var userExists UserModel
	err := collection.FindOne(context.Background(), bson.D{
		{Key: "email", Value:model.Email},
	}).Decode(&userExists)
	if err == nil && err != mongo.ErrNoDocuments {
		return core.NewHttpErr(1, 400, "user already exists")
	}
	result, err := collection.InsertOne(context.Background(), model)
	if err != nil {
		return core.NewHttpErr(2, 500, "internal server error")
	}
	stringId := result.InsertedID
	objectId, ok := stringId.(primitive.ObjectID)
	if !ok {
		return core.NewHttpErr(3, 500, "internal server error")
	}
	model.Id = objectId
	return nil
}