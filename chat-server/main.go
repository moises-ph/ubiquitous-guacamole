package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/db"
	"github.com/moises-ph/ubiquitous-guacamole/chat-server/handlers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(400, err.Error())
			return
		}
		defer conn.Close()
		for {
			conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
			time.Sleep(time.Second)
		}
	})

	router.Run() // listen and serve on 0.0.0.0:8080
}
