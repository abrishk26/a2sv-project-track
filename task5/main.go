package main

import (
	"context"
	"fmt"
	"os"

	"github.com/abrishk26/a2sv-project-track/task5/router"

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

	collection := client.Database("test").Collection("tasks")

	router := router.NewRouter(collection)

	router.Run()
}
