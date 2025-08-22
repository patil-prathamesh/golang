package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patil-prathamesh/golang-jwt-project/database"
	helper "github.com/patil-prathamesh/golang-jwt-project/helpers"
	"github.com/patil-prathamesh/golang-jwt-project/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var collection *mongo.Collection = database.OpenCollection(database.Client, "movies")
var validate = validator.New()

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var foundUser models.User

	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Email is incorrect",
		})
		return
	}

	passwordValid, msg := VerifyPassword(user.Password, foundUser.Password)
	if !passwordValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, err := helper.GenerateAllTokens(foundUser.Email, foundUser.First_name, foundUser.Last_name, foundUser.User_type, foundUser.User_id)

	fmt.Print(token, refreshToken, " ----------")

	helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)
	err = collection.FindOne(ctx, bson.M{"user_id": foundUser.User_id}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(http.StatusOK, foundUser)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":         foundUser.User_id,
			"email":      foundUser.Email,
			"first_name": foundUser.First_name,
			"last_name":  foundUser.Last_name,
			"user_type":  foundUser.User_type,
		},
		"token":         token,
		"refresh_token": refreshToken,
	})
}
func Signup(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation error"})
		return
	}
	count, err := collection.CountDocuments(ctx, bson.M{"email": user.Email})

	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking email"})
	}

	password := HashPassword(user.Password)
	user.Password = password

	count, err = collection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking phone"})
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "this email or password already exist.",
		})
		return
	}

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()

	token, refresh_token, _ := helper.GenerateAllTokens(user.Email, user.First_name, user.Last_name, user.User_type, user.User_id)

	user.Token = token
	user.Refresh_token = refresh_token

	insertion_number, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not created"})
		return
	}

	c.JSON(http.StatusCreated, insertion_number)

}
func GetUsers(c *gin.Context) {
	if err := helper.CheckUserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	recordsPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
	if err != nil || recordsPerPage < 1 {
		recordsPerPage = 10
	}
	page, err := strconv.Atoi(c.Query("page"))

	if err != nil || page < 1 {
		page = 1
	}

	startIndex := (page - 1) * recordsPerPage

	startIndex, err = strconv.Atoi(c.Query("startIndex"))
	
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{}},
	}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
			{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
		}},
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: 1},
			{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordsPerPage}}}},
		}},
	}

	result, err := collection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing user items"})
		return
	}

	var allUsers []bson.M

	if err = result.All(ctx, &allUsers); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, allUsers[0])

}

func GetUser(c *gin.Context) {
	userId := c.Param("user_id")
	fmt.Println("User id", userId, " &&&&&&&&&")
	if err := helper.MatchUserTypeToUid(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User
	err := collection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedPassword)
}

func VerifyPassword(plainPassword string, hashedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	check := true
	msg := " "
	if err != nil {
		msg = "Password is incorrect"
		check = false
	}
	return check, msg
}
