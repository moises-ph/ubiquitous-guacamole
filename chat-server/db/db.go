package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() {
	user := os.Getenv("User")
	pass := os.Getenv("Password")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://"+user+":"+pass+"@chat.0qyc9ll.mongodb.net/?retryWrites=true&w=majority&appName=chat"))
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	var Users = client.Database("Chat-server").Collection("users")

	println("Conected to the database")
}
