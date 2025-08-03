package main

import (
	"log"

	"go/mongo/models"
  "go/mongo/controllers"
  "github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	r := gin.New()
	// r.Use(gin.Logger())
	models.ConnectDatabase()

	log.Println("Server started!")

	movies := r.Group("/movies")
	{
		movies.POST("/", controllers.CreateMovie)
		movies.PUT("/:id", controllers.UpdateMovie)
		movies.DELETE("/:id", controllers.DeleteMovie)
		movies.DELETE("/", controllers.DeleteAllMovies)
		movies.GET("/", controllers.ListAllMovies)
		movies.GET("/one/:name", controllers.FindMovieByName)
		movies.GET("/all/:name", controllers.FindAllMoviesByName)
		movies.POST("/multiple", controllers.InsertMultipleMovies)

	}
	// models.InsertMovie(`{
	// "movie" : "Avengers",
	// "actors": "["Robert"]"
	// }`)
	r.Run()

}


// package main

// import (
//   "context"
//   "fmt"

//   "go.mongodb.org/mongo-driver/v2/mongo"
//   "go.mongodb.org/mongo-driver/v2/mongo/options"
//   "go.mongodb.org/mongo-driver/v2/mongo/readpref"
// )

// func main() {
//   // Use the SetServerAPIOptions() method to set the version of the Stable API on the client
//   serverAPI := options.ServerAPI(options.ServerAPIVersion1)
//   opts := options.Client().ApplyURI("mongodb+srv://prathameshpatil2906:qIF0jZAGf4xCMfvr@cluster0.bj4sw0l.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

//   // Create a new client and connect to the server
//   client, err := mongo.Connect(opts)
//   if err != nil {
//     panic(err)
//   }

//   defer func() {
//     if err = client.Disconnect(context.TODO()); err != nil {
//       panic(err)
//     }
//   }()

//   // Send a ping to confirm a successful connection
//   if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
//     panic(err)
//   }
//   fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
// }
