package models

import (
	"context"
	"htmx-cares/src/core"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionModel struct {
	Id primitive.ObjectID `bson:"_id" json:"_id"`
	User primitive.ObjectID `bson:"user" json:"user"`
	Expiration time.Time `bson:"expiration" json:"expiration"`
}

func NewSessionModel(userId primitive.ObjectID) *SessionModel {
	expiration := time.Now().Add(24 * time.Hour) // 24 hours
	return &SessionModel{
		User:       userId,
		Expiration: expiration.UTC(),
	}
}

func (m *SessionModel) ClearUserSessions(sessionCollection *mongo.Collection) (*core.HttpErr) {
	_, err := sessionCollection.DeleteMany(context.Background(), bson.D{
		{Key: "user", Value: m.User},
	})
	if err != nil {
		return core.NewHttpErr(0, 500, "internal server error")
	}
	return nil
}

func (m *SessionModel) Insert(sessionCollection *mongo.Collection) (*core.HttpErr) {
	result, err := sessionCollection.InsertOne(context.Background(), bson.D{
		{Key: "user", Value: m.User},
		{Key: "expiration", Value: m.Expiration},
	})
	if err != nil {
		return core.NewHttpErr(0, 500, "internal server error")
	}
	stringId := result.InsertedID
	objectId, ok := stringId.(primitive.ObjectID)
	if !ok {
		return core.NewHttpErr(1, 500, "internal server error")
	}
	m.Id = objectId
	return nil
}