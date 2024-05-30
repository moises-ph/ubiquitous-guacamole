package main

import (
	"github.com/gin-gonic/gin"
)

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
	/*router.POST("/post", func(c *gin.Context) {
		c.
	})*/
	router.Run() // listen and serve on 0.0.0.0:8080
}
