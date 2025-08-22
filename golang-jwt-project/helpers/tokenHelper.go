package helpers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/patil-prathamesh/golang-jwt-project/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Uid       string
	UserType  string
	jwt.StandardClaims
}

func GenerateAllTokens(email string, first_name string, last_name string, user_type string, uid string) (string, string, error) {
	fmt.Println("Hello********************")
	err := godotenv.Load(".env")
	secretKey := os.Getenv("JWT_SECRET")
	if err != nil {
		panic("Error loading .env file")
	}
	claims := &SignedDetails{
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		Uid:       uid,
		UserType:  user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		Email:     email,
		FirstName: first_name,
		LastName:  last_name,
		Uid:       uid,
		UserType:  user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	fmt.Println(token, err)
	if err != nil {
		return " ", " ", nil
	}


	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		return " ", " ", nil
	}

	return token, refreshToken, nil
}

func UpdateAllTokens(token string, refreshToken string, userId string) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	collection := database.OpenCollection(database.Client, "movies")
	filter := bson.M{"user_id": userId}
	update := bson.M{"$set": bson.M{
		"token":         token,
		"refresh_token": refreshToken,
	}}
	collection.UpdateOne(ctx, filter, update)

}

func ValidateToken(signedToken string) (SignedDetails, string) {
	godotenv.Load(".env")
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	var msg string

	if err != nil {
		msg = err.Error()
		return SignedDetails{}, msg
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "token is invalid"
		return SignedDetails{}, msg
	}

	if claims.ExpiresAt < time.Now().Unix() {
		msg = "token is expired"
		return SignedDetails{}, msg
	}

	return *claims, msg
}
