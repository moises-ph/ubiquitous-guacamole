package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/db"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.ConnectDb()
	router := gin.Default()

	//Users routes
	router.POST("/user", handlers.UserRegister)
	router.GET("/user", handlers.GetUser)

	router.Run() // listen and serve on 0.0.0.0:8080
}
