package models

import (
	"context"
	"fmt"
	"htmx-cares/src/core"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	Email string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

func NewUserModel(email string, password string) (*UserModel) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return &UserModel{
		Email: email,
		Password: string(hashedPassword),
	}
}

func (model *UserModel) Insert(collection *mongo.Collection) (*core.HttpErr) {
	var userExists UserModel
	err := collection.FindOne(context.Background(), bson.D{
		{Key: "email", Value:model.Email},
	}).Decode(&userExists)
	if err == nil && err != mongo.ErrNoDocuments {
		return core.NewHttpErr(0, 400, "user already exists")
	}
	result, err := collection.InsertOne(context.Background(), bson.D{
		{Key: "email", Value: model.Email},
		{Key: "password", Value: model.Password},
	})
	if err != nil {
		fmt.Println(err)
		return core.NewHttpErr(1, 500, "internal server error")
	}
	stringId := result.InsertedID
	objectId, ok := stringId.(primitive.ObjectID)
	if !ok {
		return core.NewHttpErr(2, 500, "internal server error")
	}
	model.Id = objectId
	return nil
}

func (model *UserModel) Find(collection *mongo.Collection) (*core.HttpErr) {
	err := collection.FindOne(context.Background(), bson.D{{
		Key: "email", Value: model.Email,
	}}).Decode(model)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return core.NewHttpErr(0, 400, "invalid credentials")
		}
		return core.NewHttpErr(1, 500, "internal server error")
	}
	return nil
}