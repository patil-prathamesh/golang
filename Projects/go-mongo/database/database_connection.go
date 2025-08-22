package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var database string

func ConnectDatabase() {
	mongoUri := os.Getenv("MONGODB_URI")
	database = os.Getenv("DATABASE")
	clientOPtions := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(context.Background(), clientOPtions)
	if err != nil {
		panic(err)
	}

	mongoClient = client
}

func GetCollection() *mongo.Collection {
	collection := os.Getenv("COLLECTION")
	return mongoClient.Database(database).Collection(collection)
}