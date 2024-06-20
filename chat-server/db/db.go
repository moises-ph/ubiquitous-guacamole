package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
			println("disconnected from db")
			panic(err)
		}
	}()

	var Users = client.Database("Chat-server").Collection("users")

	var Name string = ""

	var result bson.M
	err = Users.FindOne(context.TODO(), bson.D{{"Key", Name}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found")
		return
	}

	println("Conected to the database")
}
