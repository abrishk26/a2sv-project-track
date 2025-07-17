package router

import (
	"github.com/gin-gonic/gin"

	"github.com/abrishk26/a2sv-project-track/task5/controllers"
	"github.com/abrishk26/a2sv-project-track/task5/data"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(coll *mongo.Collection) *gin.Engine {
	controller := controllers.NewTaskController(data.NewTaskManager(coll))
	router := gin.Default()

	router.GET("/tasks/:id", controller.GetTask)
	router.GET("/tasks", controller.GetTasks)
	router.POST("/tasks", controller.CreateTask)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.DELETE("/tasks/:id", controller.DeleteTask)

	return router
}
