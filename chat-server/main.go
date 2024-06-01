package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/handlers"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/structs"
)

var users structs.User

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "some test string to send",
			"boolean": true,
		})
	})

	router.POST("/postest", handlers.UserRegister)

	router.Run() // listen and serve on 0.0.0.0:8080
}
