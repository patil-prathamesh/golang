package helpers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/patil-prathamesh/golangg-jwt/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, first_name string, last_name string, user_type string, user_id string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		First_name: first_name,
		Last_name: last_name,
		Uid: user_id,
		User_type: user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{"token": signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token": signedRefreshToken})

	Updated_at, _ := time.Now()
	updateObj = append(updateObj, bson.E{"updated_at": Updated_at})
	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOneOptions{
		Upsert: &upsert,
	}
	_, err := userCollection.UpdateOne(ctx,filter,bson.D{"$set", bson.E(updateObj[]), &opt},)

	defer cancel()
	if err != nil {
		log.Panic(err)
		return 
	}
	return
}
