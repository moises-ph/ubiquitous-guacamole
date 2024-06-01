package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/structs"
)

func UserRegister(c *gin.Context) {
	var newUser structs.User
	c.BindJSON(&newUser)

	c.JSON(200, newUser)
}
