package main

import (
	"fmt"
	"mongo/go/controller"
	"mongo/go/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// "go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/v2/bson"
// 	"go.mongodb.org/mongo-driver/v2/mongo"
// 	"go.mongodb.org/mongo-driver/v2/mongo/options"

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ENV loaded...")
	}
	models.Init()
	router := gin.New()
	movies := router.Group("/movies")
	movies.POST("/", controller.AddMovie)
	router.Run()
}