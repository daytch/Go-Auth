package db

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func GetDBCollection() (*mongo.Collection, error) {
	options.Client().ApplyURI("mongodb://Admin:Mapan1326@5.189.134.84:27017/admin")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// client, err := mongo.Connect(context.TODO(), "mongodb://Admin:Mapan1326@5.189.134.84:27017/admin")
	if err != nil {)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database("GoLogin").Collection("users")
	return collection, nil
}
