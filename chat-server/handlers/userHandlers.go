package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/structs"
)

func UserRegister(c *gin.Context) {
	var newUser structs.User

	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(400, err.Error())
	}
	c.JSON(200, newUser)
}
