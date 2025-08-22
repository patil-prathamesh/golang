package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type User struct {
//     Name  string `binding:"required"`
//     Email string `binding:"required,email"`
//     Age   int    `binding:"min=18,max=100"`
// }

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string              `json:"movie,omitempty" binding:"required"`
	Watched bool                `json:"watched,omitempty" binding:"required"`
}

// insert 1 record 

func InsertOneMovie(movie Netflix) error {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil{
		return err
	}
	fmt.Println("Inserted movie...",inserted.InsertedID, inserted.Acknowledged)
	return nil
}