package router

import (
	"github.com/gin-gonic/gin"

	"github.com/abrishk26/a2sv-project-track/task7/Delivery/controllers"
	"github.com/abrishk26/a2sv-project-track/task7/Infrastructure"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(taskColl, userColl *mongo.Collection) *gin.Engine {
	taskController := controllers.NewTaskController(taskColl)
	userController := controllers.NewUserController(userColl)

	router := gin.Default()

	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login/:id", userController.LoginUser)

	logged := router.Group("/")
	logged.Use(infrastructures.IsLoggedIn())

	{
		logged.GET("/tasks/:id", taskController.GetTask)
		logged.GET("/tasks", taskController.GetTasks)
	}

	authorized := router.Group("/")
	authorized.Use(infrastructures.IsAdmin())

	{
		authorized.POST("/tasks", taskController.CreateTask)
		authorized.PUT("/tasks/:id", taskController.UpdateTask)
		authorized.DELETE("/tasks/:id", taskController.DeleteTask)
		authorized.PUT("/users/:id", userController.UpdateUser)
		authorized.GET("/users/:id", userController.GetUser)
		authorized.GET("/users", userController.GetUsers)
		authorized.DELETE("/users/:id", userController.DeleteUser)
	}

	return router
}
