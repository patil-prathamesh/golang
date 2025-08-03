package models

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectDatabase() {
	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println(os.Getenv("COLLECTION_NAME"))
	fmt.Println(os.Getenv("MONGODB_URI"))
	connectionString := os.Getenv("MONGODB_URI")
	// db := os.Getenv("DB_NAME")
	// collName := os.Getenv("COLLECTION_NAME")
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}

	mongoClient = client
}

func GetCollection() *mongo.Collection {
	return mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
}
