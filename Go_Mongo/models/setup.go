package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://prathameshpatil2906:qIF0jZAGf4xCMfvr@cluster0.bj4sw0l.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const db = "sample_supplies"
const collName = "movies"

var mongoClient *mongo.Client

func ConnectDatabase() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}

	mongoClient = client
}