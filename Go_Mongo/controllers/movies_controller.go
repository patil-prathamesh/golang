package controllers

import (
	"go/mongo/models"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}

	err := models.InsertMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to insert movie"})
	}
	c.JSON(http.StatusCreated,gin.H{"message":"Movie inserted successfully"})
}

func UpdateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id := c.Param("id")
	err := models.UpdateMovie(id, movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to update movie"})
	}
	c.JSON(http.StatusOK, gin.H{"message":"Movie updated successfully"})
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := models.DeleteMovie(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error": "Failed to delete movie"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Movie delted successfully"})
}

func DeleteAllMovies(c *gin.Context) {
	err := models.DeleteAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all movies"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted all movies"})
}

func ListAllMovies(c *gin.Context) {
	movies := models.ListAll()
	c.JSON(http.StatusOK, movies)
}

func FindMovieByName(c *gin.Context) {
	movieName := c.Param("name")
	err, result := models.Find(movieName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
	}
	c.JSON(http.StatusOK, result)
}

func FindAllMoviesByName(c *gin.Context) {
	movieName := c.Param("name")
	movies := models.FindAll(movieName)
	if len(movies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error":"No movies found"})
	}
	c.JSON(http.StatusOK, movies)
} 

func InsertMultipleMovies(c *gin.Context) {
	var movies []models.Movie

	if err := c.ShouldBindJSON(&movies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err := models.InsertMany(movies)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert movies"})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Movies inserted successfully"})
}