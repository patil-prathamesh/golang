package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Movie struct {
	ID     primitive.ObjectID `json:"_id,omitempty"  bson:"_id,omitempty"`
	Movie  string             `json:"movie"`
	Actors []string           `jsoon:"actors"`
}

func InsertMovie(movie Movie) error {
	collection := mongoClient.Database(db).Collection(collName)
	inserted, err := collection.InsertOne(context.TODO(), movie)
	fmt.Println(inserted)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a record with id: ", inserted.InsertedID)
	return err
}

func InsertMany(movies []Movie) error {
	newMovies := make([]interface{}, len(movies))
	for i, movie := range movies {
		newMovies[i] = movie
	}
	collection := mongoClient.Database(db).Collection(collName)

	result, err := collection.InsertMany(context.TODO(), newMovies)
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
	log.Println(result)
	return err
}

func UpdateMovie(movieId string, movie Movie) error {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"movie": movie.Movie, "actors": movie.Actors}}

	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	fmt.Println(result)
	if err != nil {
		return err
	}
	fmt.Println("New record: ", result)
	return nil
}

func DeleteMovie(movieId string) error {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}

	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.DeleteOne(context.TODO(), filter)
	fmt.Println(result)
	if err != nil {
		return err
	}
	fmt.Println("Delete result: ", result)
	return nil
}

func Find(movieName string) (error, Movie) {
	var result Movie
	filter := bson.M{"movie": movieName}
	collection := mongoClient.Database(db).Collection(collName)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	fmt.Println("Hello")
	// 	log.Fatal(err)
	// }
	return err,result
}

func FindAll(movieName string) []Movie {
	var results []Movie
	filter := bson.M{"movie": movieName}
	colllection := mongoClient.Database(db).Collection(collName)
	cursor, err := colllection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cursor)
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func ListAll() []Movie {
	var results []Movie
	colllection := mongoClient.Database(db).Collection(collName)
	cursor, err := colllection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func DeleteAll() error {
	collection := mongoClient.Database(db).Collection(collName)
	delResult, err := collection.DeleteMany(context.TODO(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Records deleted... ", delResult.DeletedCount)
	return nil
}
