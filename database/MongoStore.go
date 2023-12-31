package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	Client *mongo.Client
	UserCollection *mongo.Collection
	SessionCollection *mongo.Collection
	LocationCollection *mongo.Collection
}

func NewMongoStore() (*MongoStore, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}
	userCollection := Collection(client, "users")
	sessionCollection := Collection(client, "sessions")
	locationCollection := Collection(client, "locations")
	store := &MongoStore {
		Client: client,
		UserCollection: userCollection,
		SessionCollection: sessionCollection,
		LocationCollection: locationCollection,
	}
	return store, nil
}

func Connect() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetMaxPoolSize(10))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Collection(client *mongo.Client, name string) (coll *mongo.Collection) {
	coll = client.Database(os.Getenv("DB_NAME")).Collection(name)
	return coll
}