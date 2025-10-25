package models

import (
	"context"

	"peddie-backend/library/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectDatabase() {
	clientOption := options.Client().ApplyURI(constants.ConnectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}
	mongoClient = client
}
