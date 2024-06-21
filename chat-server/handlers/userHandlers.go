package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/db"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/structs"
	"go.mongodb.org/mongo-driver/bson"
)

func UserRegister(c *gin.Context) {
	var newUser structs.User

	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(400, err.Error())
	}

	result, err := db.Users.InsertOne(context.TODO(), bson.D{
		{"Id", newUser.Id},
		{"Name", newUser.Name},
		{"Cellphone", newUser.Cellphone},
		{"Email", newUser.Email},
		{"Password", newUser.Password},
	})

	if err != nil {
		c.JSON(400, err.Error())
	}
	println(result)
	c.JSON(200, newUser)
}

type userEmail struct {
	Email string
}

func GetUser(c *gin.Context) {
	var result bson.M
	var email userEmail
	c.BindJSON(&email)
	err := db.Users.FindOne(context.TODO(), bson.D{{"Email", email.Email}}).Decode(&result)
	if err != nil {
		c.JSON(400, "user does not exist")
	}

	c.JSON(200, result)
}
