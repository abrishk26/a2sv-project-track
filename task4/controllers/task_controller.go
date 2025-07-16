package controllers

import (
	"net/http"
	"strconv"

	"github.com/abrishk26/a2sv-project-track/task4/data"
	"github.com/abrishk26/a2sv-project-track/task4/models"
	"github.com/gin-gonic/gin"
)

func NewTaskController(tm *data.TaskManager) *TaskController {
	return &TaskController{tm}
}

type TaskController struct {
	taskManager *data.TaskManager
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	task = tc.taskManager.Add(task)

	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "tasks retrieved successfully",
		"tasks": tc.taskManager.GetAll(),
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

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	task, err := tc.taskManager.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task retrieved successfully",
		"task": task,
	})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id is not specified",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	var task models.Task

	err = c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	task, err = tc.taskManager.Update(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "task updated successfully",
		"task": task,
	})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id is not specified",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	task, err := tc.taskManager.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "task deleted successfully",
		"task": task,
	})
}
