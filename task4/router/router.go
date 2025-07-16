package router

import (
	"github.com/gin-gonic/gin"

	"github.com/abrishk26/data"
	"github.com/abrishk26/controllers"
)

func NewRouter() *gin.Engine {
	controller := controllers.NewTaskController(data.NewTaskManager())
	router := gin.Default()

	Router.GET("/tasks/:id", controller.GetTask)
	Router.POST("/tasks", controller.CreateTask)
	Router.PUT("/tasks/:id", controller.UpdateTask)
	Router.DELETE("/tasks/:id", controller.DeleteTask)

	return router
}




