package router

import (
	"github.com/abrishk26/a2sv-project-track/task8/Delivery/controllers"
	"github.com/abrishk26/a2sv-project-track/task8/Infrastructure"

	"github.com/gin-gonic/gin"
)

func CreateTaskRoute(r *gin.Engine, taskController *controllers.TaskController) {
	logged := r.Group("/tasks")
	logged.Use(infrastructures.IsLoggedIn)

	{
		logged.GET("/:id", taskController.GetTask)
		logged.GET("/", taskController.GetTasks)
		logged.POST("/", taskController.CreateTask)
		logged.PUT("/:id", taskController.UpdateTask)
		logged.DELETE("/:id", taskController.DeleteTask)
	}
}

func CreateUserRoute(r *gin.Engine, userController *controllers.UserController) {
	r.POST("/users/register", userController.RegisterUser)
	r.POST("/users/login/:id", userController.LoginUser)

	logged := r.Group("/users")
	logged.Use(infrastructures.IsLoggedIn)

	{
		logged.PUT("/:id", userController.UpdateUser)
		logged.GET("/:id", userController.GetUser)
		logged.GET("/", userController.GetUsers)
		logged.DELETE("/:id", userController.DeleteUser)
	}
}
