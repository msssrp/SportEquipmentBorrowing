package database

import (
	"context"
	"github.com/msssrp/SportEquipmentBorrowing/function"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Database, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri, err := function.GetDotEnv("MONGOURL")
	if err != nil {
		return nil, err
	}
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		client.Disconnect(context.TODO())
		return nil, err
	}
	return client.Database("admin", nil), nil
}
