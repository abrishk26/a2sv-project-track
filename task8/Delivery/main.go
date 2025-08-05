package main

import (
	"context"
	"fmt"
	"os"

	"github.com/abrishk26/a2sv-project-track/task8/Delivery/controllers"
	"github.com/abrishk26/a2sv-project-track/task8/Delivery/routers"
	"github.com/abrishk26/a2sv-project-track/task8/Infrastructure"
	"github.com/abrishk26/a2sv-project-track/task8/Repositories"
	"github.com/abrishk26/a2sv-project-track/task8/Usecases"
	"github.com/gin-gonic/gin"

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

	taskRepository := repositories.NewTaskRepository(taskCollection)
	userRepository := repositories.NewUserRepository(userCollection)

	passwordService := infrastructures.NewPasswordService()
	tokenService := infrastructures.NewTokenService([]byte(os.Getenv("HASH_KEY")))

	userUsecase := usecases.NewUserUsecases(userRepository, passwordService, tokenService)
	taskUsecase := usecases.NewTaskUsecases(taskRepository, userRepository, tokenService)

	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	r := gin.Default()

	router.CreateTaskRoute(r, taskController)
	router.CreateUserRoute(r, userController)

	r.Run()
}
