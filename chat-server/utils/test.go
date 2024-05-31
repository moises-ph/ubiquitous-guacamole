package utils

import "github.com/gin-gonic/gin"

func Testing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "some test string to send in another module",
		"boolean": true,
	})
}
