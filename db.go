package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dataBase *mongo.Database

func StartMongo(uri, dbName string, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		panic(err)
	}

	// if err != nil {
	// 	return nil, err
	// }

	dataBase = client.Database(dbName)
}

func CloseMongo(db *mongo.Database) error {
	return db.Client().Disconnect(context.Background())
}
