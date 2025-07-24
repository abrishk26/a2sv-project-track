package main

import (
	"context"
	"fmt"
	"os"

	"github.com/abrishk26/a2sv-project-track/task7/Delivery/routers"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		return
	}

	clientOption := options.Client().ApplyURI(os.Getenv("mongo"))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(clientOption)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		return
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		return
	}

	fmt.Println("Database connected sucessfully")

	taskCollection := client.Database("test").Collection("tasks")
	userCollection := client.Database("test").Collection("users")

	router := router.NewRouter(taskCollection, userCollection)

	router.Run()
}
