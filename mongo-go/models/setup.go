package models

import (
	"log"
	"os"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var collection *mongo.Collection

func Init() {
	mongo_uri := os.Getenv("MONGODB_URI")
	db := os.Getenv("DATABASE_NAME")
	colName := os.Getenv("COLLECTION_NAME")
	fmt.Println(mongo_uri, db, colName)
	clientOption := options.Client().ApplyURI(mongo_uri)

	client, err := mongo.Connect(clientOption)
	if err != nil{
		log.Fatal("error", err.Error())
	}
	fmt.Println("Database Connected!!")

	collection = client.Database(db).Collection(colName)
	fmt.Println("Collection is ready")
}