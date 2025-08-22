package controller

import (
	"fmt"
	"mongo/go/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func helper(c *gin.Context, status int64, key string, value string) {
	c.JSON(int(status), gin.H{key: value})
}

func AddMovie(c *gin.Context) {
	var movie models.Netflix
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := movie.ID.Hex()
	fmt.Println(movie, "------------", id)
	movie.ID = primitive.NewObjectID()
	err := models.InsertOneMovie(movie)
	response := struct {
        ID      string `json:"id"`
        Message string `json:"message"`
        Data    interface{} `json:"data"`
    }{
        ID:      movie.ID.Hex(),
        Message: "Movie added successfully",
        Data:    movie,
    }
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert movie"})
		return
	}
	c.JSON(http.StatusOK, response)
}
