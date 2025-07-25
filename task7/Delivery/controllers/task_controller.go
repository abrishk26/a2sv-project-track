package controllers

import (
	"net/http"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"
	"github.com/abrishk26/a2sv-project-track/task7/Repositories"
	usecases "github.com/abrishk26/a2sv-project-track/task7/Usecases"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewTaskController(coll *mongo.Collection) *TaskController {
	return &TaskController{coll}
}

type TaskController struct {
	coll *mongo.Collection
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	taskID, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.ID = taskID.String()

	taskUsecases := usecases.NewTaskUsecases(repositories.NewTaskRepository(c.Request.Context(), tc.coll))
	err = taskUsecases.Add(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	taskUsecases := usecases.NewTaskUsecases(repositories.NewTaskRepository(c.Request.Context(), tc.coll))
	tasks, err := taskUsecases.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "tasks retrieved successfully",
		"tasks":   tasks,
	})
}

func (tc *TaskController) GetTask(c *gin.Context) {
	idParam := c.Param("id")

	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id missing",
		})
		return
	}

	taskUsecases := usecases.NewTaskUsecases(repositories.NewTaskRepository(c.Request.Context(), tc.coll))
	task, err := taskUsecases.Get(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task retrieved successfully",
		"task":    task,
	})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id missing",
		})
		return
	}

	var task domain.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	taskUsecases := usecases.NewTaskUsecases(repositories.NewTaskRepository(c.Request.Context(), tc.coll))
	err = taskUsecases.Update(idParam, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedTask, err := taskUsecases.Get(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "task updated successfully",
		"task":    updatedTask,
	})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id missing",
		})
		return
	}

	taskUsecases := usecases.NewTaskUsecases(repositories.NewTaskRepository(c.Request.Context(), tc.coll))
	deletedTask, err := taskUsecases.Get(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = taskUsecases.Delete(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "task deleted successfully",
		"task":    deletedTask,
	})
}
