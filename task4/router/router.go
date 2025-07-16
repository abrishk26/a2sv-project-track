package router

import (
	"github.com/gin-gonic/gin"

	"github.com/abrishk26/a2sv-project-track/task4/data"
	"github.com/abrishk26/a2sv-project-track/task4/controllers"
)

func NewRouter() *gin.Engine {
	controller := controllers.NewTaskController(data.NewTaskManager())
	router := gin.Default()

	router.GET("/tasks/:id", controller.GetTask)
	router.GET("/tasks", controller.GetTasks)
	router.POST("/tasks", controller.CreateTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.DELETE("/tasks/:id", controller.DeleteTask)

	return router
}




