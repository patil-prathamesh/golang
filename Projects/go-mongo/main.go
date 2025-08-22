package main

import (
	"fmt"
	"prthmsh-mongo/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	port := os.Getenv("PORT")

	if err != nil {
		log.Panic("Error while loading .env")
	}
	database.ConnectDatabase()
	
}
