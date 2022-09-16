package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DATABASE_URI    = "mongodb://root:root@localhost:27017/"
	DATABASE_NAME   = "blogdb"
	COLLECTION_NAME = "blog"
)

func connectToMongoDB() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI(DATABASE_URI))
	if err != nil {
		log.Printf("failed to create new mongoDB client: %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Printf("failed to connect to to mongoDB: %v\n", err)
	}

	return client.Database(DATABASE_NAME).Collection(COLLECTION_NAME)
}
