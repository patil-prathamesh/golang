package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URI")

	options := options.Client().ApplyURI(MongoDb)
	client, err := mongo.Connect(options)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collection string) *mongo.Collection {
	return Client.Database("sample_supplies").Collection(collection)
}
