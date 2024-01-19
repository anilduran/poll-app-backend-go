package main

import (
	"example.com/poll-app-backend-go/db"
	"example.com/poll-app-backend-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("failed to load .env file!")
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	err = db.InitializeDB()

	if err != nil {
		panic("failed to connect to the db!")
	}

	r.Run(":8080")

}
